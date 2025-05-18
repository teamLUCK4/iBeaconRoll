package main

import (
	"fmt"
	"log"

	"iBeaconRoll-server/config"
	"iBeaconRoll-server/database"
	"iBeaconRoll-server/routes"
)

func main() {
	// 설정 로드
	cfg := config.LoadConfig()
	
	// 데이터베이스 연결
	database.InitDB(cfg.DatabaseURL)
	defer database.CloseDB()
	
	// 라우터 설정
	router := routes.SetupRouter()
	
	// 서버 시작
	serverAddr := ":" + cfg.Port
	fmt.Printf("서버가 %s 포트에서 실행 중입니다\n", cfg.Port)
	log.Fatal(router.Run(serverAddr))
}