package redis

import (
	goredis "github.com/redis/go-redis/v9"
)

// RedisGateway provides an interface for interacting with Redis for ranking and user progress management.
type RedisGateway struct {
	client   *goredis.Client
	redisKey string
}

// NewRedisGateway initializes a new RedisGateway instance.
func NewRedisGateway(client *goredis.Client, redisKey string) *RedisGateway {
	return &RedisGateway{
		client:   client,
		redisKey: redisKey,
	}
}
