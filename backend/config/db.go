// 📄 backend/config/db.go

package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"

    _ "github.com/lib/pq" // PostgreSQL 드라이버 등록
)

var DB *sql.DB

func ConnectDB() {
    dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        os.Getenv("DB_HOST"),
        os.Getenv("DB_PORT"),
        os.Getenv("DB_USER"),
        os.Getenv("DB_PASSWORD"),
        os.Getenv("DB_NAME"),
    )

    var err error
    DB, err = sql.Open("postgres", dsn)
    if err != nil {
        log.Fatal("❌ DB 연결 실패:", err)
    }

    if err = DB.Ping(); err != nil {
        log.Fatal("❌ DB Ping 실패:", err)
    }

    log.Println("✅ DB 연결 완료")
}
