package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetRankResponse struct {
	Players []PlayerPosition `json:"players"`
}

type PlayerPosition struct {
	UserID   string `json:"user_id"`
	Position int    `json:"position"`
	Score    int    `json:"score"`
}

func getLeaderboard(ctx context.Context) (GetRankResponse, error) {
	url := fmt.Sprintf("%s/rank/", apiBaseURL)

	resp, err := http.Get(url)
	if err != nil {
		return GetRankResponse{}, err
	}
	defer resp.Body.Close()

	var result GetRankResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return GetRankResponse{}, err
	}

	return result, nil
}
