package controllers

import (
	"net/http"
	"time"

	"iBeaconRoll-server/config"
	"iBeaconRoll-server/models"

	"github.com/gin-gonic/gin"
)

func UpdateAttendance(c *gin.Context) {
	var req models.AttendanceUpdateRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 날짜 유효성 검증 (선택)
	_, err := time.Parse("2006-01-02", req.AttendanceDate)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid date format. Use YYYY-MM-DD."})
		return
	}

	db := config.PostgresDB

	// 먼저 해당 레코드가 있는지 확인
	var exists bool
	checkQuery := `
		SELECT EXISTS (
			SELECT 1 FROM attendances 
			WHERE student_id = $1 AND timetable_id = $2 AND attendance_date = $3
		)
	`
	err = db.Get(&exists, checkQuery, req.StudentID, req.TimetableID, req.AttendanceDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to check attendance record"})
		return
	}

	if !exists {
		// 레코드가 없으면 INSERT
		insertQuery := `
			INSERT INTO attendances (student_id, timetable_id, classroom, attendance_date, status, attendance_time)
			VALUES ($1, $2, $3, $4, $5, CURRENT_TIME)
		`
		_, err = db.Exec(insertQuery, req.StudentID, req.TimetableID, req.Classroom, req.AttendanceDate, req.Status)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to insert attendance"})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "Attendance inserted",
			"action":  "insert",
		})
		return
	}

	// 레코드가 있으면 UPDATE
	updateQuery := `
		UPDATE attendances 
		SET status = $1, attendance_time = CURRENT_TIME
		WHERE student_id = $2 AND timetable_id = $3 AND attendance_date = $4
	`
	result, err := db.Exec(updateQuery, req.Status, req.StudentID, req.TimetableID, req.AttendanceDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attendance"})
		return
	}

	rows, _ := result.RowsAffected()
	c.JSON(http.StatusOK, gin.H{
		"message":       "Attendance updated",
		"action":        "update",
		"rows_affected": rows,
	})
}
