package questions

import (
	"context"
	"errors"
	"fmt"

	"quiz/domain/entities"
)

var (
	ErrInvalidAnswer = errors.New("answer is invalid")
)

// AnswerQuestionInput Input structure for submitting an answer.
type AnswerQuestionInput struct {
	UserID     string
	QuestionID int
	Answer     int
}

// AnswerQuestionOutput Output structure for the answer submission result.
type AnswerQuestionOutput struct {
	IsCorrect bool
}

// AnswerQuestionUseCase Use case for submitting an answer.
type AnswerQuestionUseCase struct {
	redis  AnswerQuestionRedisGateway
	memory AnswerQuestionMemoryGateway
}

// AnswerQuestionRedisGateway Interface defining methods for Redis interactions in answer submission.
type AnswerQuestionRedisGateway interface {
	IncrementScore(ctx context.Context, userID string) error
	SetUserProgress(ctx context.Context, userID string, questionID int) error
}

// AnswerQuestionMemoryGateway Interface defining methods for retrieving questions in answer submission.
type AnswerQuestionMemoryGateway interface {
	GetQuestion(ctx context.Context, questionID int) (entities.Question, error)
}

// NewAnswerQuestionUseCase Constructor for AnswerQuestionUseCase.
func NewAnswerQuestionUseCase(redis AnswerQuestionRedisGateway, memory AnswerQuestionMemoryGateway) *AnswerQuestionUseCase {
	return &AnswerQuestionUseCase{redis: redis, memory: memory}
}

// AnswerQuestion Submits an answer and updates the user's progress.
func (g *AnswerQuestionUseCase) AnswerQuestion(ctx context.Context, input AnswerQuestionInput) (AnswerQuestionOutput, error) {
	if input.Answer > 2 {
		return AnswerQuestionOutput{}, fmt.Errorf("validation error: %w", ErrInvalidAnswer)
	}

	question, err := g.memory.GetQuestion(ctx, input.QuestionID)
	if err != nil {
		return AnswerQuestionOutput{}, fmt.Errorf("failed to get question: %w", err)
	}

	if err = g.redis.SetUserProgress(ctx, input.UserID, input.QuestionID+1); err != nil {
		return AnswerQuestionOutput{}, fmt.Errorf("failed to set user progress: %w", err)
	}

	if input.Answer != question.AnswerID {
		return AnswerQuestionOutput{IsCorrect: false}, nil
	}

	if err = g.redis.IncrementScore(ctx, input.UserID); err != nil {
		return AnswerQuestionOutput{}, fmt.Errorf("failed to set score: %w", err)
	}

	return AnswerQuestionOutput{IsCorrect: true}, nil
}
