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

type FullError struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

//go:generate moq -stub -pkg mocks -out mocks/answer_question_use_case.go . AnswerQuestionUseCase

// AnswerQuestionRequest Request structure for submitting an answer.
type AnswerQuestionRequest struct {
	NumberQuestion int `json:"number_question" validate:"required"`
	Answer         int `json:"answer" validate:"required"`
}

// AnswerQuestionResponse Response structure for submitting an answer.
type AnswerQuestionResponse struct {
	IsCorrect bool `json:"is_correct"`
}

// AnswerQuestionUseCase Interface defining methods for submitting an answer.
type AnswerQuestionUseCase interface {
	AnswerQuestion(ctx context.Context, input questions.AnswerQuestionInput) (questions.AnswerQuestionOutput, error)
}

// AnswerQuestion
// @Summary Submit an Answer
// @Description Allows a user to submit an answer to a question
// @Tags Question
// @Accept json
// @Produce json
// @Param user query string true "User ID"
// @Param requestBody body AnswerQuestionRequest true "User answer submission"
// @Success 200 {object} AnswerQuestionResponse
// @Failure 400 {object} pkg.FullError
// @Failure 500 {object} pkg.FullError
// @Router /question/ [post]
func (h Handler) AnswerQuestion(w http.ResponseWriter, r *http.Request) {
	const operation = "Question.QuestionHandler.AnswerQuestion"
	ctx := r.Context()

	l := pkg.FromCtx(ctx).With(
		zap.String("operation", operation),
		zap.String("path", r.RequestURI),
		zap.String("method", r.Method),
	)

	userID := r.URL.Query().Get("user")
	if userID == "" {
		l.Error("'user' parameter is missing")
		responseError := pkg.ErrBadRequest.WithTitle("'user' parameter is missing")
		_ = pkg.Send(w, responseError, http.StatusBadRequest)
		return
	}

	var requestBody AnswerQuestionRequest
	err := json.NewDecoder(r.Body).Decode(&requestBody)
	if err != nil {
		l.Error("failed to get player position", zap.String("user_id", userID), zap.Error(err))
		responseError := pkg.ErrBadRequest.WithTitle("Body is empty or has no valid fields")
		_ = pkg.Send(w, responseError, http.StatusBadRequest)
		return
	}

	output, err := h.submitAnswerUseCase.AnswerQuestion(ctx, questions.AnswerQuestionInput{
		UserID:     userID,
		Answer:     requestBody.Answer,
		QuestionID: requestBody.NumberQuestion,
	})

	if err != nil {
		var fullError pkg.FullError
		statusCode := http.StatusBadRequest

		switch {
		case errors.Is(err, questions.ErrInvalidAnswer):
			fullError = pkg.ErrInavlidAnswer
		default:
			fullError = pkg.ErrInternalServerError
			statusCode = http.StatusInternalServerError
			l.Error("failed to submit answer", zap.Error(err))
		}
		_ = pkg.Send(w, fullError, statusCode)
		return
	}

	response := AnswerQuestionResponse{IsCorrect: output.IsCorrect}
	_ = pkg.Send(w, response, http.StatusOK)
}
