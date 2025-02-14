package questions_test

import (
	"context"
	"testing"

	redis2 "github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"quiz/domain/usecases/questions"
	"quiz/infrastructure/gateways/redis"
	repository "quiz/infrastructure/memory"
	"quiz/pkg/testutils"
)

func TestAnswerQuestion_Success(t *testing.T) {
	tests := []struct {
		name          string
		input         questions.AnswerQuestionInput
		expectCorrect bool
		expectScore   float64
	}{
		{
			name: "should be successful when the user answers correctly",
			input: questions.AnswerQuestionInput{
				UserID:     "joana",
				QuestionID: 1,
				Answer:     0, // Correct answer
			},
			expectCorrect: true,
			expectScore:   2,
		},
		{
			name: "should be successful when the user answers incorrectly",
			input: questions.AnswerQuestionInput{
				UserID:     "joana",
				QuestionID: 1,
				Answer:     1, // Incorrect answer
			},
			expectCorrect: false,
			expectScore:   1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ctx := context.Background()
			redisClient, questionStorage, redisGateway := setupTestEnv(t, ctx)
			defer redisClient.Close()

			useCase := questions.NewAnswerQuestionUseCase(redisGateway, questionStorage)

			// submit answer
			answerOutput, err := useCase.AnswerQuestion(ctx, tt.input)
			require.NoError(t, err)
			assert.Equal(t, tt.expectCorrect, answerOutput.IsCorrect)

			// Assert score update in Redis
			playerPosition, err := redisGateway.GetPlayerPosition(ctx, tt.input.UserID)
			require.NoError(t, err)
			assert.Equal(t, tt.expectScore, playerPosition.Score)

			// Assert the answer stored is the expected one
			question, err := questionStorage.GetQuestion(ctx, tt.input.QuestionID)
			require.NoError(t, err)
			if tt.expectCorrect {
				assert.Equal(t, tt.input.Answer, question.AnswerID)
			} else {
				assert.NotEqual(t, tt.input.Answer, question.AnswerID)
			}
		})
	}
}

func TestAnswerQuestion_Failure(t *testing.T) {
	ctx := context.Background()
	redisClient, questionStorage, redisGateway := setupTestEnv(t, ctx)
	defer redisClient.Close()

	useCase := questions.NewAnswerQuestionUseCase(redisGateway, questionStorage)

	input := questions.AnswerQuestionInput{
		UserID:     "joana",
		QuestionID: 1,
		Answer:     5, // Invalid answer (out of range)
	}

	// Submit answer
	answerOutput, err := useCase.AnswerQuestion(ctx, input)

	// Assert error
	require.Error(t, err)
	assert.ErrorIs(t, err, questions.ErrInvalidAnswer)

	// Assert empty response on failure
	assert.Empty(t, answerOutput)
}

// setupTestEnv initializes the test environment with Redis and memory storage.
func setupTestEnv(t *testing.T, ctx context.Context) (*redis2.Client, *repository.MemoryQuestionStorage, *redis.RedisGateway) {
	redisClient := testutils.NewRedisClient(t, ctx)

	// Flush all Redis data before each test to prevent data persistence issues
	err := redisClient.FlushAll(ctx).Err()
	require.NoError(t, err)

	questionStorage := repository.NewMemoryQuestionStorage()
	redisGateway := redis.NewRedisGateway(redisClient, "leaderboard")

	// Initialize a user in the leaderboard to avoid errors when retrieving player position
	err = redisGateway.IncrementScore(ctx, "joana")
	require.NoError(t, err)

	return redisClient, questionStorage, redisGateway
}
