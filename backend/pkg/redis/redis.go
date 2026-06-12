package redispkg

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v8"

	"github.com/Iichengyu1207/nexus-I/backend/internal/config"
)

// NewClient 创建 Redis 客户端连接（无 Redis 时返回 nil，业务代码判空降级）
func NewClient(cfg *config.Config) (*redis.Client, error) {
	if cfg.RedisHost == "" {
		return nil, nil
	}

	addr := fmt.Sprintf("%s:%d", cfg.RedisHost, cfg.RedisPort)
	client := redis.NewClient(&redis.Options{
		Addr: addr,
		DB:   cfg.RedisDB,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
