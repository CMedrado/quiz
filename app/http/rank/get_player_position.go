package rank

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	"quiz/domain/usecases/rank"
	"quiz/pkg"
)

//go:generate moq -stub -pkg mocks -out mocks/get_player_position_use_case.go . GetPlayerPositionUseCase

// GetPlayerPositionResponse Response structure for getting a player's position.
type GetPlayerPositionResponse struct {
	Player PlayerPosition `json:"player"`
}

// GetPlayerPositionUseCase Interface defining methods for getting a player's position.
type GetPlayerPositionUseCase interface {
	GetPlayerPosition(ctx context.Context, input rank.GetPlayerPositionInput) (rank.GetPlayerPositionOutput, error)
}

// GetPlayerPosition
// @Summary Get Player Position
// @Description Fetches the ranking position of a player
// @Tags Rank
// @Produce json
// @Param user query string true "User ID"
// @Success 200 {object} GetPlayerPositionResponse
// @Failure 400 {object} pkg.FullError
// @Failure 500 {object} pkg.FullError
// @Router /rank/player/ [get]
func (h Handler) GetPlayerPosition(w http.ResponseWriter, r *http.Request) {
	const operation = "Rank.RankHandler.GetPlayerPosition"
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

	output, err := h.getPlayerPositionUseCase.GetPlayerPosition(ctx, rank.GetPlayerPositionInput{UserID: userID})
	if err != nil {
		l.Error("failed to get player position", zap.String("user_id", userID), zap.Error(err))
		_ = pkg.Send(w, pkg.ErrInternalServerError, http.StatusInternalServerError)
		return
	}

	response := GetPlayerPositionResponse{
		Player: PlayerPosition{
			UserID:   output.PlayerPosition.UserID,
			Position: output.PlayerPosition.Position,
			Score:    output.PlayerPosition.Score,
		},
	}

	_ = pkg.Send(w, response, http.StatusOK)
}
