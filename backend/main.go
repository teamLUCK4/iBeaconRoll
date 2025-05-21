// backend/main.go
package main

import (
	"fmt"
	"iBeaconRoll-server/config"
	"iBeaconRoll-server/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 iBeaconRoll server started!")

	// 1. PostgreSQL 연결
	config.InitPostgres()

	// 2. Gin 서버 초기화
	r := gin.Default()

	// 3. API 라우트 등록
	routes.RegisterAttendanceRoutes(r)

	// 4. 서버 실행
	fmt.Println("🚀 서버 실행 중: http://localhost:8080")
	r.Run(":8080")

}
