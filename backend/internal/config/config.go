package config

import (
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	AppEnv     string
	DBType     string
	SQLitePath string
	RedisHost  string
	RedisPort  int
	RedisDB    int
	JWTSecret  string
}

func Load() *Config {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "development"
	}

	viper.SetConfigName(".env." + env)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.AddConfigPath("../") // backend 子目录运行时向上找
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		// 没有 .env 文件时继续使用系统环境变量或默认值
	}

	return &Config{
		AppEnv:     viper.GetString("APP_ENV"),
		DBType:     viper.GetString("DB_TYPE"),
		SQLitePath: viper.GetString("SQLITE_PATH"),
		RedisHost:  viper.GetString("REDIS_HOST"),
		RedisPort:  viper.GetInt("REDIS_PORT"),
		RedisDB:    viper.GetInt("REDIS_DB"),
		JWTSecret:  viper.GetString("JWT_SECRET"),
	}
}
