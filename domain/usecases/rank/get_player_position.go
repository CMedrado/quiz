package rank

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"

	"quiz/domain/entities"
)

// GetPlayerPositionInput Input structure for retrieving a player's position.
type GetPlayerPositionInput struct {
	UserID string
}

// GetPlayerPositionOutput Output structure containing the player's position details.
type GetPlayerPositionOutput struct {
	PlayerPosition entities.PlayerPosition
}

// GetPlayerPositionUseCase Use case for fetching a player's position from Redis.
type GetPlayerPositionUseCase struct {
	redis GetPlayerPositionRedisGateway
}

// GetPlayerPositionRedisGateway Interface defining the method for retrieving player position from Redis.
type GetPlayerPositionRedisGateway interface {
	GetPlayerPosition(ctx context.Context, ID string) (goredis.RankScore, error)
}

// NewGetPlayerPositionUseCase Constructor for RetrievePlayerPositionUseCase.
func NewGetPlayerPositionUseCase(redis GetPlayerPositionRedisGateway) *GetPlayerPositionUseCase {
	return &GetPlayerPositionUseCase{redis: redis}
}

// GetPlayerPosition Retrieves the player's position in the ranking.
func (g *GetPlayerPositionUseCase) GetPlayerPosition(ctx context.Context, input GetPlayerPositionInput) (GetPlayerPositionOutput, error) {
	playerPositionRankScore, err := g.redis.GetPlayerPosition(ctx, input.UserID)
	if err != nil {
		return GetPlayerPositionOutput{}, fmt.Errorf("failed to retrieve player position: %w", err)
	}

	playerPosition := entities.PlayerPosition{
		Position: int(playerPositionRankScore.Rank + 1),
		UserID:   input.UserID,
		Score:    int(playerPositionRankScore.Score),
	}

	return GetPlayerPositionOutput{PlayerPosition: playerPosition}, nil
}
