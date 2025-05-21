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

	query := `
		UPDATE attendances 
		SET status = $1 
		WHERE student_id = $2 AND classroom = $3 AND attendance_date = $4
	`

	result, err := db.Exec(query, req.Status, req.StudentID, req.Classroom, req.AttendanceDate)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update attendance"})
		return
	}

	rows, _ := result.RowsAffected()
	c.JSON(http.StatusOK, gin.H{
		"message":       "Attendance updated",
		"rows_affected": rows,
	})
}
