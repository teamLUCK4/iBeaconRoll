package routes

import (
	"github.com/gin-gonic/gin"

	"iBeaconRoll-server/controllers"
)

func RegisterScheduleRoutes(router *gin.Engine) {
	router.GET("/api/students/:id/schedule/today", controllers.GetStudentTodaySchedule)
}
