// 📄 backend/config/db.go

package config

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // PostgreSQL 드라이버 등록

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var PostgresDB *sqlx.DB
var MongoDB *mongo.Database

func InitPostgres() {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
	)

	var err error
	PostgresDB, err = sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatal("❌ DB 연결 실패:", err)
	}

	if err = PostgresDB.Ping(); err != nil {
		log.Fatal("❌ DB Ping 실패:", err)
	}

	log.Println("✅ PostgreSQL 연결 완료")
}

func InitMongoDB() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// MongoDB 연결 설정
	mongoURI := fmt.Sprintf("mongodb://%s:%s",
		os.Getenv("MONGO_HOST"),
		os.Getenv("MONGO_PORT"),
	)

	clientOptions := options.Client().ApplyURI(mongoURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal("❌ MongoDB 연결 실패:", err)
	}

	// 연결 확인
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("❌ MongoDB Ping 실패:", err)
	}

	MongoDB = client.Database(os.Getenv("MONGO_DB_NAME"))
	log.Println("✅ MongoDB 연결 완료")
}
