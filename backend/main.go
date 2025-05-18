// backend/main.go
package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"iBeaconRoll-server/config"
	"iBeaconRoll-server/routes"
)

func main() {
	fmt.Println("🚀 iBeaconRoll server started!")

	// 1. 설정 로드
	cfg := config.LoadConfig()
	
	// 2. 데이터베이스 연결
	config.InitDB()
	defer config.CloseDB()
	
	// 3. Gin 서버 초기화
	router := gin.Default()

	// 4. API 라우트 등록
	routes.RegisterAttendanceRoutes(router)
	
	// 5. 서버 실행
	serverAddr := ":" + cfg.Port
	fmt.Printf("🚀 서버 실행 중: http://localhost%s\n", serverAddr)
	log.Fatal(router.Run(serverAddr))
}