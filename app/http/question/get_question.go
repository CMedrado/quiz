package question

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"go.uber.org/zap"

	"quiz/domain/usecases/questions"
	"quiz/pkg"
)

//go:generate moq -stub -pkg mocks -out mocks/get_question_use_case.go . GetQuestionUseCase

// GetQuestionRequest Request structure for getting a question.
type GetQuestionRequest struct {
	UserID string `json:"user_id" validate:"required"`
}

// GetQuestionResponse Response structure for getting a question.
type GetQuestionResponse struct {
	NumberQuestion int      `json:"number_question"`
	Question       string   `json:"question"`
	Answers        []string `json:"answers"`
}

// GetQuestionUseCase Interface defining methods for getting a question.
type GetQuestionUseCase interface {
	GetQuestion(ctx context.Context, input questions.GetQuestionInput) (questions.GetQuestionOutput, error)
}

// GetQuestion
// @Summary Get a Question
// @Description Fetches the next available question for a user
// @Tags Question
// @Produce json
// @Success 200 {object} GetQuestionResponse
// @Failure 400 {object} pkg.FullError
// @Failure 500 {object} pkg.FullError
// @Router /question/ [get]
func (h Handler) GetQuestion(w http.ResponseWriter, r *http.Request) {
	const operation = "Question.QuestionHandler.GetQuestion"
	ctx := r.Context()

	l := pkg.FromCtx(ctx).With(
		zap.String("operation", operation),
		zap.String("path", r.RequestURI),
		zap.String("method", r.Method),
	)

	var requestBody GetQuestionRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		l.Error("failed to get player position", zap.Error(err))
		_ = pkg.Send(w, pkg.ErrBadRequest.WithTitle("Body is empty or has no valid fields"), http.StatusBadRequest)
		return
	}

	output, err := h.getQuestionUseCase.GetQuestion(ctx, questions.GetQuestionInput{UserID: requestBody.UserID})
	if err != nil {
		var fullError pkg.FullError
		statusCode := http.StatusBadRequest

		switch {
		case errors.Is(err, questions.ErrEmptyUserID):
			fullError = pkg.ErrInavlidUserID
		default:
			fullError = pkg.ErrInternalServerError
			statusCode = http.StatusInternalServerError
			l.Error("failed to get question", zap.String("user_id", requestBody.UserID), zap.Error(err))
		}
		_ = pkg.Send(w, fullError, statusCode)
		return
	}

	response := GetQuestionResponse{
		NumberQuestion: output.Question.ID,
		Question:       output.Question.Question,
		Answers:        output.Question.Answer,
	}
	_ = pkg.Send(w, response, http.StatusOK)
}
