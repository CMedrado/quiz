package rank_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"quiz/domain/usecases/rank"
)

func TestRank_GetPlayerPosition_Success(t *testing.T) {
	t.Run("should return correct player position when user exists in the leaderboard", func(t *testing.T) {
		// Prepare test environment
		ctx := context.Background()
		redisClient, redisGateway := setupLeaderboardTestEnv(t, ctx)
		defer redisClient.Close()

		// Simulate leaderboard data in Redis
		userID := "joana"
		err := redisGateway.IncrementScore(ctx, userID)
		require.NoError(t, err)

		// Ensure Redis is correctly storing the score
		playerPosition, err := redisGateway.GetPlayerPosition(ctx, userID)
		require.NoError(t, err)

		useCase := rank.NewGetPlayerPositionUseCase(redisGateway)

		// Call the use case
		output, err := useCase.GetPlayerPosition(ctx, rank.GetPlayerPositionInput{UserID: userID})

		// Assertions
		require.NoError(t, err)
		assert.Equal(t, playerPosition.Score, float64(output.PlayerPosition.Score))
		assert.Equal(t, userID, output.PlayerPosition.UserID)
		assert.Equal(t, playerPosition.Rank+1, int64(output.PlayerPosition.Position))
	})
}

func TestRank_GetPlayerPosition_Failure(t *testing.T) {
	tests := []struct {
		name          string
		userID        string
		shouldCloseDB bool
	}{
		{
			name:          "should return error when redis fails",
			userID:        "joana",
			shouldCloseDB: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare test environment
			ctx := context.Background()
			redisClient, redisGateway := setupLeaderboardTestEnv(t, ctx)
			redisClient.Close() // Simulating Redis failure

			useCase := rank.NewGetPlayerPositionUseCase(redisGateway)

			// Call the use case
			output, err := useCase.GetPlayerPosition(ctx, rank.GetPlayerPositionInput{UserID: tt.userID})

			// Assertions
			assert.Error(t, err)
			assert.Empty(t, output)
		})
	}
}
