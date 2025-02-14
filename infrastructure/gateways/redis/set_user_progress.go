package redis

import (
	"context"
	"fmt"
)

// SetUserProgress updates the user's progress with the given question ID.
func (r *RedisGateway) SetUserProgress(ctx context.Context, userID string, questionID int) error {
	if err := r.client.Set(ctx, userID, questionID, 0).Err(); err != nil {
		return fmt.Errorf("failed to set user progress: %w", err)
	}
	return nil
}
