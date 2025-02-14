package redis

import (
	"context"
	"fmt"
)

// IncrementScore increments the player's score by 1 in the sorted set.
func (r *RedisGateway) IncrementScore(ctx context.Context, userID string) error {
	if err := r.client.ZIncrBy(ctx, r.redisKey, 1, userID).Err(); err != nil {
		return fmt.Errorf("failed to update score: %w", err)
	}
	return nil
}
