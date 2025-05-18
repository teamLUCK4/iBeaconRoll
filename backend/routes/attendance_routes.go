package routes

import (
	"github.com/gin-gonic/gin"
	"iBeaconRoll-server/controllers"
)

func RegisterAttendanceRoutes(router *gin.Engine) {
	router.PUT("/api/attendance", controllers.UpdateAttendance)
}
