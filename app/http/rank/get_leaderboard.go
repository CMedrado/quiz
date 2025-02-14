package rank

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"quiz/domain/usecases/rank"
	"quiz/pkg"
)

//go:generate moq -stub -pkg mocks -out mocks/get_leaderboard_use_case.go . GetLeaderboardUseCase

// GetLeaderboardResponse Response structure for ranking retrieval.
type GetLeaderboardResponse struct {
	Players []PlayerPosition `json:"players"`
}

// GetLeaderboardUseCase Interface defining methods for getting a rank.
type GetLeaderboardUseCase interface {
	GetLeaderboard(ctx context.Context) (rank.GetLeaderboardOutput, error)
}

// GetLeaderboard
// @Summary Get Leaderboard
// @Description Retrieves the overall leaderboard
// @Tags Rank
// @Produce json
// @Success 200 {object} GetLeaderboardResponse
// @Failure 400 {object} pkg.FullError
// @Failure 500 {object} pkg.FullError
// @Router /rank/ [get]
func (h Handler) GetLeaderboard(w http.ResponseWriter, r *http.Request) {
	const operation = "Rank.RankHandler.GetLeaderboard"
	ctx := r.Context()

	l := pkg.FromCtx(ctx).With(
		zap.String("operation", operation),
		zap.String("path", r.RequestURI),
		zap.String("method", r.Method),
	)

	output, err := h.getRankUseCase.GetLeaderboard(ctx)
	if err != nil {
		var fullError pkg.FullError
		statusCode := http.StatusInternalServerError

		l.Error("failed to get leaderboard", zap.Error(err))
		fullError = pkg.ErrInternalServerError
		_ = pkg.Send(w, fullError, statusCode)
		return
	}

	var playersResponse GetLeaderboardResponse
	for _, entityPlayer := range output.Leaderboard {
		player := PlayerPosition{
			UserID:   entityPlayer.UserID,
			Position: entityPlayer.Position,
			Score:    entityPlayer.Score,
		}
		playersResponse.Players = append(playersResponse.Players, player)
	}

	_ = pkg.Send(w, playersResponse, http.StatusOK)
}
