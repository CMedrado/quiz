package rank

import (
	"context"
	"fmt"

	goredis "github.com/redis/go-redis/v9"

	"quiz/domain/entities"
)

// GetLeaderboardOutput Output structure for retrieving the ranking list.
type GetLeaderboardOutput struct {
	Leaderboard []entities.PlayerPosition
}

// GetLeaderboardUseCase Use case for fetching the ranking list from Redis.
type GetLeaderboardUseCase struct {
	redis GetLeaderboardRedisGateway
}

// GetLeaderboardRedisGateway Interface defining the method for retrieving the ranking from Redis.
type GetLeaderboardRedisGateway interface {
	GetLeaderboard(ctx context.Context) ([]goredis.Z, error)
}

// NewGetLeaderboardUseCase Constructor for GetLeaderboardUseCase.
func NewGetLeaderboardUseCase(redis GetLeaderboardRedisGateway) *GetLeaderboardUseCase {
	return &GetLeaderboardUseCase{redis: redis}
}

// GetLeaderboard Retrieves the ranking list.
func (g *GetLeaderboardUseCase) GetLeaderboard(ctx context.Context) (GetLeaderboardOutput, error) {
	rankZ, err := g.redis.GetLeaderboard(ctx)
	if err != nil {
		return GetLeaderboardOutput{}, fmt.Errorf("failed to retrieve ranking: %w", err)
	}

	leaderboard := make([]entities.PlayerPosition, 0, len(rankZ))
	for i, player := range rankZ {
		userID, ok := player.Member.(string)
		if !ok {
			return GetLeaderboardOutput{}, fmt.Errorf("failed to parse userID")
		}

		playerPosition := entities.PlayerPosition{
			Position: i + 1,
			UserID:   userID,
			Score:    int(player.Score),
		}
		leaderboard = append(leaderboard, playerPosition)
	}

	return GetLeaderboardOutput{Leaderboard: leaderboard}, nil
}
