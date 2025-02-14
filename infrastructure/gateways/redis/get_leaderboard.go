package redis

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

const (
	minRank = "0"
	maxRank = "10"
)

// GetLeaderboard retrieves the top 10 ranked players from the sorted set.
func (r *RedisGateway) GetLeaderboard(ctx context.Context) ([]goredis.Z, error) {
	result, err := r.client.ZRevRangeByScoreWithScores(ctx, r.redisKey, &goredis.ZRangeBy{Min: minRank, Max: maxRank}).Result()
	if err != nil {
		return nil, fmt.Errorf("failed to retrieve ranking: %w", err)
	}
	return result, nil
}
