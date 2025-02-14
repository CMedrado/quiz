package redis

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"
)

// GetPlayerPosition retrieves the rank and score of a specific player.
func (r *RedisGateway) GetPlayerPosition(ctx context.Context, ID string) (goredis.RankScore, error) {
	result, err := r.client.ZRevRankWithScore(ctx, r.redisKey, ID).Result()
	if err != nil {
		return goredis.RankScore{}, fmt.Errorf("failed to retrieve player position: %w", err)
	}
	return result, nil
}
