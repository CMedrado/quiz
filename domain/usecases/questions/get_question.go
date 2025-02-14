package questions

import (
	"context"
	"errors"
	"fmt"

	"quiz/domain/entities"
	repository "quiz/infrastructure/memory"
)

//go:generate moq -stub -pkg mocks -out mocks/get_player_position_use_case.go . GetQuestionUseCase
//go:generate moq -stub -pkg mocks -out mocks/get_player_position_use_case.go . GetQuestionRedisGateway

var (
	ErrEmptyUserID = errors.New("user id cannot be empty")
)

// GetQuestionInput Input structure for retrieving a question.
type GetQuestionInput struct {
	UserID string
}

// GetQuestionOutput Output structure containing the retrieved question.
type GetQuestionOutput struct {
	Question entities.Question
}

// GetQuestionUseCase Use case for fetching a question.
type GetQuestionUseCase struct {
	memory GetQuestionMemoryGateway
	redis  GetQuestionRedisGateway
}

// GetQuestionRedisGateway Interface defining methods to get user progress from Redis.
type GetQuestionRedisGateway interface {
	GetPlayerProgress(ctx context.Context, userID string) (int, error)
}

// GetQuestionMemoryGateway Interface defining methods to retrieve a question from memory.
type GetQuestionMemoryGateway interface {
	GetQuestion(ctx context.Context, userProgress int) (entities.Question, error)
}

// NewGetQuestionUseCase Constructor for GetQuestionUseCase.
func NewGetQuestionUseCase(memory GetQuestionMemoryGateway, redis GetQuestionRedisGateway) *GetQuestionUseCase {
	return &GetQuestionUseCase{memory: memory, redis: redis}
}

// GetQuestion Retrieves the next question for the user.
func (g *GetQuestionUseCase) GetQuestion(ctx context.Context, input GetQuestionInput) (GetQuestionOutput, error) {
	if input.UserID == "" {
		return GetQuestionOutput{}, ErrEmptyUserID
	}

	userProgress, err := g.redis.GetPlayerProgress(ctx, input.UserID)
	if err != nil {
		return GetQuestionOutput{}, fmt.Errorf("failed to retrieve user progress: %w", err)
	}

	question, err := g.memory.GetQuestion(ctx, userProgress)
	if err != nil {
		if errors.Is(err, repository.ErrRepositoryNoMoreQuestions) {
			return GetQuestionOutput{}, nil
		}
		return GetQuestionOutput{}, fmt.Errorf("failed to retrieve question: %w", err)
	}

	return GetQuestionOutput{Question: question}, nil
}
