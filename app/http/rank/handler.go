package rank

// PlayerPosition holds player ranking details.
type PlayerPosition struct {
	UserID   string `json:"user_id"`
	Position int    `json:"position"`
	Score    int    `json:"score"`
}

// Handler struct for rank-related HTTP requests.
type Handler struct {
	getRankUseCase           GetLeaderboardUseCase
	getPlayerPositionUseCase GetPlayerPositionUseCase
}

// NewRankHandler Constructor for RankHandler.
func NewRankHandler(
	getRankUseCase GetLeaderboardUseCase,
	getPlayerPositionUseCase GetPlayerPositionUseCase,
) *Handler {
	return &Handler{
		getRankUseCase:           getRankUseCase,
		getPlayerPositionUseCase: getPlayerPositionUseCase,
	}
}
