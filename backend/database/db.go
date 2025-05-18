package database

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL 드라이버
)

// DB는 데이터베이스 연결을 저장합니다.
var DB *sqlx.DB

// InitDB는 데이터베이스 연결을 초기화합니다.
func InitDB(dataSourceName string) {
	var err error
	DB, err = sqlx.Connect("postgres", dataSourceName)
	if err != nil {
		log.Fatalf("데이터베이스 연결 실패: %v", err)
	}

	if err = DB.Ping(); err != nil {
		log.Fatalf("데이터베이스 핑 실패: %v", err)
	}

	fmt.Println("데이터베이스 연결 성공")
}

// CloseDB는 데이터베이스 연결을 닫습니다.
func CloseDB() {
	if DB != nil {
		DB.Close()
	}
}