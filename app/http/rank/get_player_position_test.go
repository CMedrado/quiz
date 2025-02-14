package rank_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"go.uber.org/zap"

	"quiz/app/http/rank"
	"quiz/app/http/rank/mocks"
	"quiz/domain/entities"
	usecasesRank "quiz/domain/usecases/rank"
	"quiz/pkg"
)

func TestHandler_GetPlayerPosition_Success(t *testing.T) {
	t.Parallel()
	userID := "Zequinha"
	getPlayerResponse, getPlayerOutput := getPlayerPositionMock(userID)
	wantBody, _ := json.Marshal(getPlayerOutput)

	// Mock the dependency
	handler := rank.NewRankHandler(
		&mocks.GetLeaderboardUseCaseMock{},
		&mocks.GetPlayerPositionUseCaseMock{
			GetPlayerPositionFunc: func(_ context.Context, _ usecasesRank.GetPlayerPositionInput) (usecasesRank.GetPlayerPositionOutput, error) {
				return getPlayerResponse, nil
			},
		},
	)

	// Create request with query param "user"
	request := httptest.NewRequest(http.MethodGet, "/rank/player?user="+userID, nil)
	request = request.WithContext(pkg.WithCtx(context.Background(), zap.NewExample()))
	response := httptest.NewRecorder()

	// Execute handler
	http.HandlerFunc(handler.GetPlayerPosition).ServeHTTP(response, request)

	// Assertions
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
	assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
}

func TestHandler_GetPlayerPosition_Failure(t *testing.T) {
	t.Parallel()
	tests := []struct {
		name     string
		useCase  rank.GetPlayerPositionUseCase
		request  string
		wantCode int
		wantBody interface{}
	}{
		{
			name: "should return 400 and error when id is empty",
			useCase: &mocks.GetPlayerPositionUseCaseMock{
				GetPlayerPositionFunc: func(_ context.Context, _ usecasesRank.GetPlayerPositionInput) (usecasesRank.GetPlayerPositionOutput, error) {
					return usecasesRank.GetPlayerPositionOutput{}, nil
				},
			},
			request:  "",
			wantCode: http.StatusBadRequest,
			wantBody: pkg.ErrBadRequest.WithTitle("'user' parameter is missing"),
		},
		{
			name: "should return 500 and error when usecase error is not a domain error",
			useCase: &mocks.GetPlayerPositionUseCaseMock{
				GetPlayerPositionFunc: func(_ context.Context, _ usecasesRank.GetPlayerPositionInput) (usecasesRank.GetPlayerPositionOutput, error) {
					return usecasesRank.GetPlayerPositionOutput{}, assert.AnError
				},
			},
			request:  "florzinha",
			wantCode: http.StatusInternalServerError,
			wantBody: pkg.ErrInternalServerError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// Mock the dependency simulating an error when retrieving the player
			handler := rank.NewRankHandler(&mocks.GetLeaderboardUseCaseMock{}, tt.useCase)

			// Create request with query param "user"
			request := httptest.NewRequest(http.MethodGet, "/rank/player?user="+tt.request, nil)
			request = request.WithContext(pkg.WithCtx(context.Background(), zap.NewExample()))
			response := httptest.NewRecorder()

			wantBody, _ := json.Marshal(tt.wantBody)

			// Execute handler
			http.HandlerFunc(handler.GetPlayerPosition).ServeHTTP(response, request)

			// Assertions
			assert.Equal(t, tt.wantCode, response.Code)
			assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
			assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
		})
	}
}

// Mock function to simulate the success response
func getPlayerPositionMock(userID string) (usecasesRank.GetPlayerPositionOutput, rank.GetPlayerPositionResponse) {
	response := rank.GetPlayerPositionResponse{
		Player: rank.PlayerPosition{
			UserID:   userID,
			Position: 2,
			Score:    15,
		},
	}

	output := usecasesRank.GetPlayerPositionOutput{
		PlayerPosition: entities.PlayerPosition{
			UserID:   userID,
			Position: 2,
			Score:    15,
		},
	}

	return output, response
}
