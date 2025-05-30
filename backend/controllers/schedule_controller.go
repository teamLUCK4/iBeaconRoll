package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"iBeaconRoll-server/config"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"

	"iBeaconRoll-server/models"
)

// GetStudentTodaySchedule은 특정 학생의 오늘 시간표를 반환합니다.
// 1. postgres에서 학생의 오늘 요일 시간표 조회 -- getTodaySchedule
// 2. mongodb에서 각 교실에 해당하는 비콘 정보 조회
// 3. 시간표와 비콘 정보를 함께 반환

func getTodaySchedule(studentID int, today time.Time, dayOfWeek string) ([]models.TimetableWithAttendance, error) {
	var schedules []models.TimetableWithAttendance

	query := `
		SELECT 
			t.id, 
			t.student_id, 
			t.semester, 
			t.subject_name, 
			t.day_of_week, 
			to_char(t.start_time, 'HH24:MI:SS') as start_time, 
			to_char(t.end_time, 'HH24:MI:SS') as end_time, 
			t.classroom,
			a.status,
			to_char(a.attendance_time, 'HH24:MI:SS') as attendance_time,
			COALESCE(a.status, 'waiting') as attendance_status
		FROM 
			timetables t
		LEFT JOIN 
			attendances a ON t.id = a.timetable_id AND a.attendance_date = $1
		WHERE 
			t.student_id = $2 AND t.day_of_week = $3
		ORDER BY 
			t.start_time
	`

	err := config.PostgresDB.Select(&schedules, query, today.Format("2006-01-02"), studentID, dayOfWeek)
	if err != nil {
		log.Printf("시간표 조회 오류: %v", err)
		return nil, fmt.Errorf("시간표를 조회할 수 없습니다: %v", err)
	}

	schedulesJSON, _ := json.MarshalIndent(schedules, "", "  ")
	log.Printf("📅 조회된 시간표 데이터: %s", string(schedulesJSON))

	return schedules, nil
}

// #2. mongodb에서 각 교실에 해당하는 비콘 정보 조회
func getBeaconInfo(classroom string) (*models.Beacon, error) {
	var beaconInfo models.Beacon

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	beaconCollection := config.MongoDB.Collection("uuid")
	err := beaconCollection.FindOne(ctx, bson.M{"classroom": classroom}).Decode(&beaconInfo)
	if err != nil {
		log.Printf("📝 비콘 정보 없음 - 교실: %s", classroom)
		return nil, nil // 비콘 정보가 없는 경우 nil 반환
	}

	beaconJSON, _ := json.MarshalIndent(beaconInfo, "", "  ")
	log.Printf("📡 교실 %s의 비콘 정보: %s", classroom, string(beaconJSON))

	return &beaconInfo, nil
}

func GetStudentTodaySchedule(c *gin.Context) {
	// 학생 ID 파싱
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil || studentID <= 0 {
		log.Printf("❌ 잘못된 학생 ID: %s", c.Param("id"))
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 학생 ID입니다"})
		return
	}

	// 오늘 요일 가져오기 (Mon, Tue, Wed, Thu, Fri 형식으로 변환)
	today := time.Now()
	dayOfWeek := today.Weekday().String()[:3]

	log.Printf("📆 요청 정보 - 학생ID: %d, 날짜: %s, 요일: %s", studentID, today.Format("2006-01-02"), dayOfWeek)

	// // 주말인 경우 처리
	// if dayOfWeek == "Sat" || dayOfWeek == "Sun" {
	// 	log.Printf("💤 주말 요청 - 학생ID: %d", studentID)
	// 	c.JSON(http.StatusOK, gin.H{
	// 		"message":    "오늘은 주말입니다. 수업이 없습니다.",
	// 		"date":       today.Format("2006-01-02"),
	// 		"student_id": studentID,
	// 		"classes":    []interface{}{},
	// 	})
	// 	return
	// }

	// ## -----1. 시간표 조회----- ##
	schedules, err := getTodaySchedule(studentID, today, dayOfWeek)
	if err != nil {
		log.Printf("❌ 시간표 조회 실패: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// 시간표가 없는 경우
	if len(schedules) == 0 {
		log.Printf("📝 수업 없음 - 학생ID: %d, 날짜: %s", studentID, today.Format("2006-01-02"))
		c.JSON(http.StatusOK, gin.H{
			"message":    fmt.Sprintf("오늘(%s)은 수업이 없습니다.", dayOfWeek),
			"date":       today.Format("2006-01-02"),
			"student_id": studentID,
			"classes":    []interface{}{},
		})
		return
	}

	// ## ----- 2. 각 교실의 비콘 정보 조회 및 결합 ----- ##
	var schedulesWithBeacon []models.TimetableWithBeacon
	for _, schedule := range schedules {
		beaconInfo, err := getBeaconInfo(schedule.Classroom)
		scheduleWithBeacon := models.TimetableWithBeacon{
			TimetableWithAttendance: schedule,
			BeaconInfo:              nil,
		}

		if err == nil {
			scheduleWithBeacon.BeaconInfo = beaconInfo
		}

		schedulesWithBeacon = append(schedulesWithBeacon, scheduleWithBeacon)
	}

	// 응답 구성
	response := struct {
		Date      string                       `json:"date"`
		StudentID int                          `json:"student_id"`
		DayOfWeek string                       `json:"day_of_week"`
		Classes   []models.TimetableWithBeacon `json:"classes"`
		UpdatedAt string                       `json:"updated_at"`
	}{
		Date:      today.Format("2006-01-02T15:04:05.000Z"),
		StudentID: studentID,
		DayOfWeek: dayOfWeek,
		Classes:   schedulesWithBeacon,
		UpdatedAt: time.Now().Format("2006-01-02T15:04:05.000Z"),
	}

	responseJSON, _ := json.MarshalIndent(response, "", "  ")
	log.Printf("✅ 최종 응답 데이터: %s", string(responseJSON))

	c.JSON(http.StatusOK, response)
}
