package routes

import (
	"iBeaconRoll-server/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterAttendanceRoutes(router *gin.Engine) {
	router.PUT("/api/attendance", controllers.UpdateAttendance)
}
