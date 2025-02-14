package question_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"quiz/app/http/question"
	"quiz/app/http/question/mocks"
	"quiz/domain/usecases/questions"
	"quiz/pkg"
)

func TestHandler_AnswerQuestion_Success(t *testing.T) {
	t.Parallel()

	answerResponse, answerOutput := answerQuestionMock(true)
	wantBody, _ := json.Marshal(answerResponse)

	// Mock the dependency
	handler := question.NewQuestionHandler(&mocks.GetQuestionUseCaseMock{}, &mocks.AnswerQuestionUseCaseMock{
		AnswerQuestionFunc: func(_ context.Context, _ questions.AnswerQuestionInput) (questions.AnswerQuestionOutput, error) {
			return answerOutput, nil
		},
	})

	requestBody, _ := json.Marshal(question.AnswerQuestionRequest{
		NumberQuestion: 1,
		Answer:         2,
	})
	request := httptest.NewRequest(http.MethodPost, "/question?user=Zequinha", bytes.NewReader(requestBody))
	request.Header.Set("Content-Type", JsonContentType)
	request = request.WithContext(pkg.WithCtx(context.Background(), zap.NewExample()))
	response := httptest.NewRecorder()

	// Execute handler
	http.HandlerFunc(handler.AnswerQuestion).ServeHTTP(response, request)

	// Assertions
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
	assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
}

func TestHandler_AnswerQuestion_Failure(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		useCase  question.AnswerQuestionUseCase
		request  question.AnswerQuestionRequest
		query    string
		wantCode int
		wantBody interface{}
	}{
		{
			name: "should return 400 when answer is invalid",
			useCase: &mocks.AnswerQuestionUseCaseMock{
				AnswerQuestionFunc: func(_ context.Context, _ questions.AnswerQuestionInput) (questions.AnswerQuestionOutput, error) {
					return questions.AnswerQuestionOutput{}, questions.ErrInvalidAnswer
				},
			},
			request:  question.AnswerQuestionRequest{},
			query:    "carolzinha",
			wantCode: http.StatusBadRequest,
			wantBody: pkg.ErrInavlidAnswer,
		},

		{
			name: "should return 400 when user parameter is missing",
			useCase: &mocks.AnswerQuestionUseCaseMock{
				AnswerQuestionFunc: func(_ context.Context, _ questions.AnswerQuestionInput) (questions.AnswerQuestionOutput, error) {
					return questions.AnswerQuestionOutput{}, nil
				},
			},
			request:  question.AnswerQuestionRequest{},
			query:    "",
			wantCode: http.StatusBadRequest,
			wantBody: pkg.ErrBadRequest.WithTitle("'user' parameter is missing"),
		},
		{
			name: "should return 500 when an unexpected error occurs",
			useCase: &mocks.AnswerQuestionUseCaseMock{
				AnswerQuestionFunc: func(_ context.Context, _ questions.AnswerQuestionInput) (questions.AnswerQuestionOutput, error) {
					return questions.AnswerQuestionOutput{}, assert.AnError
				},
			},
			request:  question.AnswerQuestionRequest{NumberQuestion: 1, Answer: 2},
			query:    "carolzinha",
			wantCode: http.StatusInternalServerError,
			wantBody: pkg.ErrInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the dependency
			handler := question.NewQuestionHandler(&mocks.GetQuestionUseCaseMock{}, tt.useCase)

			requestBody, _ := json.Marshal(tt.request)
			request := httptest.NewRequest(http.MethodPost, "/question?user="+tt.query, bytes.NewReader(requestBody))
			request.Header.Set("Content-Type", JsonContentType)
			request = request.WithContext(pkg.WithCtx(context.Background(), zap.NewExample()))
			response := httptest.NewRecorder()

			wantBody, _ := json.Marshal(tt.wantBody)

			// Execute handler
			http.HandlerFunc(handler.AnswerQuestion).ServeHTTP(response, request)

			// Assertions
			assert.Equal(t, tt.wantCode, response.Code)
			assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
			assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
		})
	}
}

func answerQuestionMock(isCorrect bool) (question.AnswerQuestionResponse, questions.AnswerQuestionOutput) {
	response := question.AnswerQuestionResponse{
		IsCorrect: isCorrect,
	}

	output := questions.AnswerQuestionOutput{
		IsCorrect: isCorrect,
	}

	return response, output
}
