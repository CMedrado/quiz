package questions_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"quiz/domain/usecases/questions"
)

func TestQuestion_GetQuestion_Success(t *testing.T) {
	tests := []struct {
		name          string
		userID        string
		mockProgress  int
		expectedIndex int
	}{
		{
			name:          "should return the first question for a new user",
			userID:        "joana",
			mockProgress:  0,
			expectedIndex: 0,
		},
		{
			name:          "should return the next question based on user progress",
			userID:        "zequinha",
			mockProgress:  2, // Simulating that the user has already answered two questions
			expectedIndex: 2,
		},
		{
			name:          "should return no question when user has answered all available questions",
			userID:        "paulinho",
			mockProgress:  26, // Simulating user completed all questions
			expectedIndex: -1, // Expected to return an empty output
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare the test environment
			ctx := context.Background()
			redisClient, questionStorage, redisGateway := setupTestEnv(t, ctx)
			defer redisClient.Close()

			// Simulate user progress in Redis
			err := redisGateway.SetUserProgress(ctx, tt.userID, tt.mockProgress)
			require.NoError(t, err)

			useCase := questions.NewGetQuestionUseCase(questionStorage, redisGateway)

			// Act: Call the use case
			output, err := useCase.GetQuestion(ctx, questions.GetQuestionInput{UserID: tt.userID})

			// Assertions
			if tt.expectedIndex == -1 {
				assert.NoError(t, err) // No error, just no questions left
				assert.Empty(t, output.Question)
			} else {
				require.NoError(t, err)
				assert.Equal(t, tt.expectedIndex, output.Question.ID)
			}
		})
	}
}

func TestQuestion_GetQuestion_Failure(t *testing.T) {
	tests := []struct {
		name          string
		userID        string
		expectedError error
	}{
		{
			name:          "should return error when userID is empty",
			userID:        "",
			expectedError: questions.ErrEmptyUserID,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Prepare the test environment
			ctx := context.Background()
			redisClient, questionStorage, redisGateway := setupTestEnv(t, ctx)
			defer redisClient.Close()

			useCase := questions.NewGetQuestionUseCase(questionStorage, redisGateway)

			// Act: Call the use case
			output, err := useCase.GetQuestion(ctx, questions.GetQuestionInput{UserID: tt.userID})

			// Assertions: Ensure expected errors are returned
			assert.ErrorIs(t, err, tt.expectedError)
			assert.Empty(t, output)
		})
	}
}
