// backend/main.go
package main

import (
	"fmt"
	"iBeaconRoll-server/config"
	"iBeaconRoll-server/routes"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("🚀 iBeaconRoll server started!")

	// 1. PostgreSQL 연결
	config.InitPostgres()

	// 2. Gin 서버 초기화
	r := gin.Default()

	// 3. API 라우트 등록
	log.Println("🛣️  라우트 등록 시작...")
	routes.RegisterAttendanceRoutes(r)
	routes.RegisterScheduleRoutes(r)
	log.Println("✅ 라우트 등록 완료")

	// 등록된 라우트 확인
	for _, route := range r.Routes() {
		log.Printf("📍 Route: %s %s", route.Method, route.Path)
	}

	// 4. 서버 실행
	fmt.Println("🚀 서버 실행 중: http://localhost:8080")
	r.Run(":8080")
}
