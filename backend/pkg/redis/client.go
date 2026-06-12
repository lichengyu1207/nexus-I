package redis

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// NewClient 创建 Redis 客户端连接
func NewClient(addr, password string, db int) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: password,
		DB:       db,
	})

	if err := client.Ping(context.Background()).Err(); err != nil {
		return nil, err
	}

	return client, nil
}
