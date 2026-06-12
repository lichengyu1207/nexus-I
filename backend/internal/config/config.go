package config

import (
	"os"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	AppEnv        string
	AppPort       string
	DBDriver      string
	DBDSN         string
	RedisAddr     string
	RedisPassword string
	RedisDB       int
	JWTSecret     string
	JWTExpire     string
	LogLevel      string
}

func Load() (*Config, error) {
	_ = godotenv.Load(".env.development", ".env")

	cfg := &Config{
		AppEnv:        getEnv("APP_ENV", "development"),
		AppPort:       getEnv("APP_PORT", "8080"),
		DBDriver:      getEnv("DB_DRIVER", "sqlite"),
		DBDSN:         getEnv("DB_DSN", "./data/nexus.db"),
		RedisAddr:     getEnv("REDIS_ADDR", "localhost:6379"),
		RedisPassword: getEnv("REDIS_PASSWORD", ""),
		RedisDB:       0,
		JWTSecret:     getEnv("JWT_SECRET", "dev-secret"),
		JWTExpire:     getEnv("JWT_EXPIRE", "24h"),
		LogLevel:      getEnv("LOG_LEVEL", "debug"),
	}

	return cfg, nil
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return strings.TrimSpace(value)
	}
	return fallback
}
