package config

import (
	"os"
)

// Config 구조체는 애플리케이션 설정을 저장합니다.
type Config struct {
	DatabaseURL string
	Port        string
}

// LoadConfig 함수는 환경 변수로부터 설정을 로드합니다.
func LoadConfig() *Config {
	return &Config{
		// PostgreSQL 연결 문자열
		DatabaseURL: getEnv("DATABASE_URL", "postgresql://user:password@postgres:5432/school_db?sslmode=disable"),
		Port:        getEnv("PORT", "8080"),
	}
}

// getEnv 함수는 환경변수를 가져오고 없으면 기본값을 반환합니다.
func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}