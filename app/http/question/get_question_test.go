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
	"quiz/domain/entities"
	"quiz/domain/usecases/questions"
	"quiz/pkg"
)

const (
	JsonContentType = "application/json"
)

func TestHandler_GetQuestion_Success(t *testing.T) {
	t.Parallel()
	getQustionResponse, getQuestionOutput := getQuestionMock()
	wantBody, _ := json.Marshal(getQustionResponse)

	// Mock the dependency
	handler := question.NewQuestionHandler(&mocks.GetQuestionUseCaseMock{
		GetQuestionFunc: func(_ context.Context, _ questions.GetQuestionInput) (questions.GetQuestionOutput, error) {
			return getQuestionOutput, nil
		},
	},
		&mocks.AnswerQuestionUseCaseMock{})

	// Create request
	requestBody, _ := json.Marshal(question.GetQuestionRequest{UserID: "florzinha"})
	request := httptest.NewRequest(http.MethodGet, "/question", bytes.NewReader(requestBody))
	request = request.WithContext(pkg.WithCtx(context.Background(), zap.NewExample()))
	response := httptest.NewRecorder()

	// Execute handler
	http.HandlerFunc(handler.GetQuestion).ServeHTTP(response, request)

	// Assertions
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
	assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
}

func TestHandler_GetQuestion_Failure(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		useCase  question.GetQuestionUseCase
		request  question.GetQuestionRequest
		wantCode int
		wantBody interface{}
	}{
		{
			name: "should return 400 and error when id is empty",
			useCase: &mocks.GetQuestionUseCaseMock{
				GetQuestionFunc: func(_ context.Context, _ questions.GetQuestionInput) (questions.GetQuestionOutput, error) {
					return questions.GetQuestionOutput{}, questions.ErrEmptyUserID
				},
			},
			request:  question.GetQuestionRequest{},
			wantCode: http.StatusBadRequest,
			wantBody: pkg.ErrInavlidUserID,
		},
		{
			name: "should return 500 and error when usecase error is not a domain error",
			useCase: &mocks.GetQuestionUseCaseMock{
				GetQuestionFunc: func(_ context.Context, _ questions.GetQuestionInput) (questions.GetQuestionOutput, error) {
					return questions.GetQuestionOutput{}, assert.AnError
				},
			},
			request:  question.GetQuestionRequest{UserID: "florzinha"},
			wantCode: http.StatusInternalServerError,
			wantBody: pkg.ErrInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the dependency simulating an error when retrieving the player
			handler := question.NewQuestionHandler(tt.useCase, &mocks.AnswerQuestionUseCaseMock{})
			requestBody, _ := json.Marshal(tt.request)

			// Create request with query param "user"
			request := httptest.NewRequest(http.MethodGet, "/question", bytes.NewReader(requestBody))
			request = request.WithContext(pkg.WithCtx(context.Background(), zap.NewExample()))
			response := httptest.NewRecorder()

			wantBody, _ := json.Marshal(tt.wantBody)

			// Execute handler
			http.HandlerFunc(handler.GetQuestion).ServeHTTP(response, request)

			// Assertions
			assert.Equal(t, tt.wantCode, response.Code)
			assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
			assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
		})
	}
}

func getQuestionMock() (question.GetQuestionResponse, questions.GetQuestionOutput) {
	response := question.GetQuestionResponse{
		NumberQuestion: 1,
		Question:       "What is the national language of Malta?",
		Answers:        []string{"Italian", "Maltese", "English"},
	}

	output := questions.GetQuestionOutput{
		Question: entities.Question{
			ID:       1,
			Question: "What is the national language of Malta?",
			Answer:   []string{"Italian", "Maltese", "English"},
		},
	}

	return response, output
}
