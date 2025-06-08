// backend/main.go
package main

import (
	"fmt"
	"os"
	"iBeaconRoll-server/config"
	"iBeaconRoll-server/routes"
	"log"
	"github.com/joho/godotenv"
	"github.com/gin-gonic/gin"
)

func main() {

	// ✅ .env 파일 불러오기
	if err := godotenv.Load(); err != nil {
		log.Println("⚠️ .env 파일을 불러올 수 없습니다. 시스템 환경변수를 사용합니다.")
	} else {
		log.Println("✅ .env 파일을 성공적으로 로드했습니다.")
	}
	
	fmt.Println("🚀 iBeaconRoll server started!")

	// 1. PostgreSQL 연결
	config.InitPostgres()

	// 2. MongoDB 연결
	config.InitMongoDB()

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

	// 간단한 "Hello, World!" 엔드포인트 추가
	r.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello, Luck4!",
		})
	})

	// 4. 서버 실행
	fmt.Println("🚀 서버 실행 중: http://localhost:8080")
	r.Run(":8080")
}
