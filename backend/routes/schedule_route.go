package routes

import (
	"github.com/gin-gonic/gin"
	
	"iBeaconRoll-server/controllers"
)

// SetupRouter는 API 라우터를 설정합니다.
func SetupRouter() *gin.Engine {
	router := gin.Default()

	// CORS 설정
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Authorization")
		
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		
		c.Next()
	})

	// API 경로 그룹
	api := router.Group("/api")
	{
		students := api.Group("/students")
		{
			// 특정 학생의 오늘 시간표 조회
			students.GET("/:id/schedule/today", controllers.GetStudentTodaySchedule)
		}
	}

	return router
}