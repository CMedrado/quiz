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

const (
	JsonContentType = "application/json"
)

func TestHandler_GetLeaderboard_Success(t *testing.T) {
	t.Parallel()
	getLeaderboardOutput, getLeaderboardResponse := getLeaderboardMock()
	wantBody, _ := json.Marshal(getLeaderboardResponse)

	// Mock the dependency
	handler := rank.NewRankHandler(&mocks.GetLeaderboardUseCaseMock{
		GetLeaderboardFunc: func(ctx context.Context) (usecasesRank.GetLeaderboardOutput, error) {
			return getLeaderboardOutput, nil
		},
	}, &mocks.GetPlayerPositionUseCaseMock{})

	// Create request
	request := httptest.NewRequest(http.MethodGet, "/rank", nil)
	request = request.WithContext(pkg.WithCtx(context.Background(), zap.NewExample()))
	response := httptest.NewRecorder()

	// Execute handler
	http.HandlerFunc(handler.GetLeaderboard).ServeHTTP(response, request)

	// Assertions
	assert.Equal(t, http.StatusOK, response.Code)
	assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
	assert.Equal(t, JsonContentType, response.Header().Get("Content-Type"))
}

func TestHandler_GetLeaderboard_Failure(t *testing.T) {
	t.Parallel()
	// Mock the dependency simulating an error when retrieving the player
	handler := rank.NewRankHandler(&mocks.GetLeaderboardUseCaseMock{
		GetLeaderboardFunc: func(ctx context.Context) (usecasesRank.GetLeaderboardOutput, error) {
			return usecasesRank.GetLeaderboardOutput{}, assert.AnError
		},
	}, &mocks.GetPlayerPositionUseCaseMock{})

	// Create request
	request := httptest.NewRequest(http.MethodGet, "/rank", nil)
	request = request.WithContext(pkg.WithCtx(context.Background(), zap.NewExample()))
	response := httptest.NewRecorder()

	wantBody, _ := json.Marshal(pkg.ErrInternalServerError)

	// Execute handler
	http.HandlerFunc(handler.GetLeaderboard).ServeHTTP(response, request)

	// Assertions
	assert.Equal(t, http.StatusInternalServerError, response.Code)
	assert.Equal(t, string(wantBody), strings.TrimSpace(response.Body.String()))
	assert.Equal(t, "application/json", response.Header().Get("Content-Type"))
}

func getLeaderboardMock() (usecasesRank.GetLeaderboardOutput, rank.GetLeaderboardResponse) {
	response := rank.GetLeaderboardResponse{
		Players: []rank.PlayerPosition{
			{
				UserID:   "Luisinho",
				Position: 1,
				Score:    20,
			},
			{
				UserID:   "Zequinha",
				Position: 2,
				Score:    15,
			},
			{
				UserID:   "PatoDonald",
				Position: 3,
				Score:    10,
			},
		},
	}

	output := usecasesRank.GetLeaderboardOutput{
		Leaderboard: []entities.PlayerPosition{
			{
				UserID:   "Luisinho",
				Position: 1,
				Score:    20,
			},
			{
				UserID:   "Zequinha",
				Position: 2,
				Score:    15,
			},
			{
				UserID:   "PatoDonald",
				Position: 3,
				Score:    10,
			},
		},
	}

	return output, response
}
