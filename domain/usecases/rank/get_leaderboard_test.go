package rank_test

import (
	"context"
	"testing"

	redis2 "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"quiz/domain/entities"
	"quiz/domain/usecases/rank"
	"quiz/infrastructure/gateways/redis"
	"quiz/pkg/testutils"
)

func TestGetLeaderboard_Success(t *testing.T) {
	tests := []struct {
		name          string
		want          []entities.PlayerPosition
		expectedCount int
	}{
		{
			name: "should return leaderboard with multiple players",
			want: []entities.PlayerPosition{
				{UserID: "joana", Position: 1, Score: 3},
				{UserID: "zequinha", Position: 2, Score: 2},
				{UserID: "paulinho", Position: 3, Score: 1},
			},
			expectedCount: 3,
		},
		{
			name:          "should return empty leaderboard when no players exist",
			want:          []entities.PlayerPosition{},
			expectedCount: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare the test environment
			ctx := context.Background()
			redisClient, redisGateway := setupLeaderboardTestEnv(t, ctx)
			defer redisClient.Close()

			// Simulate leaderboard data in Redis
			for _, entry := range tt.want {
				for i := 0; i < entry.Score; i++ {
					err := redisGateway.IncrementScore(ctx, entry.UserID)
					require.NoError(t, err)
				}
			}

			useCase := rank.NewGetLeaderboardUseCase(redisGateway)

			// Call the use case
			output, err := useCase.GetLeaderboard(ctx)

			// Assertions
			require.NoError(t, err)
			assert.Len(t, output.Leaderboard, tt.expectedCount)

			// Ensure correct data mapping
			for i, expected := range tt.want {
				assert.Equal(t, expected.UserID, output.Leaderboard[i].UserID)
				assert.Equal(t, expected.Position, output.Leaderboard[i].Position)
				assert.Equal(t, expected.Score, output.Leaderboard[i].Score)
			}
		})
	}
}

func TestGetLeaderboard_Failure(t *testing.T) {
	t.Run("should return error when redis retrieval fails", func(t *testing.T) {
		// Prepare the test environment
		ctx := context.Background()
		redisClient, redisGateway := setupLeaderboardTestEnv(t, ctx)

		redisClient.Close() // Simulating Redis failure

		useCase := rank.NewGetLeaderboardUseCase(redisGateway)

		// Call the use case
		output, err := useCase.GetLeaderboard(ctx)

		// Assertions
		require.Error(t, err)
		assert.Empty(t, output.Leaderboard)
	})
}

// setupLeaderboardTestEnv initializes Redis for leaderboard testing.
func setupLeaderboardTestEnv(t *testing.T, ctx context.Context) (*redis2.Client, *redis.RedisGateway) {
	redisClient := testutils.NewRedisClient(t, ctx)
	redisGateway := redis.NewRedisGateway(redisClient, "leaderboard")

	// Flush all Redis data before each test to prevent data persistence issues
	err := redisClient.FlushAll(ctx).Err()
	require.NoError(t, err)

	return redisClient, redisGateway
}
