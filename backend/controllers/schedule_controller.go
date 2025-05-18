package controllers

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	
	"iBeaconRoll-server/database"
	"iBeaconRoll-server/models"
)

// GetStudentTodaySchedule은 특정 학생의 오늘 시간표를 반환합니다.
func GetStudentTodaySchedule(c *gin.Context) {
	// 학생 ID 파싱
	studentID, err := strconv.Atoi(c.Param("id"))
	if err != nil || studentID <= 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "유효하지 않은 학생 ID입니다"})
		return
	}

	// 오늘 요일 가져오기 (Mon, Tue, Wed, Thu, Fri 형식으로 변환)
	today := time.Now()
	dayOfWeek := today.Weekday().String()[:3]
	
	// 주말인 경우 처리
	if dayOfWeek == "Sat" || dayOfWeek == "Sun" {
		c.JSON(http.StatusOK, gin.H{
			"message": "오늘은 주말입니다. 수업이 없습니다.",
			"date": today.Format("2006-01-02"),
			"student_id": studentID,
			"classes": []interface{}{},
		})
		return
	}

	// 데이터베이스에서 학생의 오늘 요일 시간표 조회
	var schedules []models.TimetableWithAttendance
	query := `
		SELECT 
			t.id, 
			t.student_id, 
			t.semester, 
			t.subject_name, 
			t.day_of_week, 
			t.start_time, 
			t.end_time, 
			t.classroom,
			a.id AS attendance_id,
			a.status,
			a.attendance_time
		FROM 
			timetables t
		LEFT JOIN 
			attendances a ON t.id = a.timetable_id AND a.attendance_date = $1
		WHERE 
			t.student_id = $2 AND t.day_of_week = $3
		ORDER BY 
			t.start_time
	`
	err = database.DB.Select(&schedules, query, today.Format("2006-01-02"), studentID, dayOfWeek)
	
	if err != nil {
		log.Printf("시간표 조회 오류: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "시간표를 조회할 수 없습니다"})
		return
	}

	// 응답 구성
	response := models.DailySchedule{
		Date:       today,
		StudentID:  studentID,
		DayOfWeek:  dayOfWeek,
		Classes:    schedules,
		UpdatedAt:  time.Now(),
	}

	c.JSON(http.StatusOK, response)
}