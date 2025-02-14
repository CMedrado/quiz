package redis

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	goredis "github.com/redis/go-redis/v9"
)

// GetPlayerProgress retrieves the stored progress of a user.
func (r *RedisGateway) GetPlayerProgress(ctx context.Context, userID string) (int, error) {
	progress, err := r.client.Get(ctx, userID).Result()
	if err != nil {
		if errors.Is(err, goredis.Nil) {
			return 0, nil // No progress stored yet
		}
		return 0, fmt.Errorf("failed to retrieve user progress: %w", err)
	}

	progressInt, err := strconv.Atoi(progress)
	if err != nil {
		return 0, fmt.Errorf("failed to convert progress to int: %w", err)
	}

	return progressInt, nil
}
