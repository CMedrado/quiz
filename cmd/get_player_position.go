package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type getPlayerPositionResponse struct {
	Player PlayerPosition `json:"player"`
}

func getPlayerPosition(ctx context.Context, username string) (PlayerPosition, error) {
	url := fmt.Sprintf("%s/rank/player/?user=%s", apiBaseURL, username)

	resp, err := http.Get(url) // Faz a requisição HTTP GET
	if err != nil {
		return PlayerPosition{}, err
	}
	defer resp.Body.Close()

	var result getPlayerPositionResponse
	err = json.NewDecoder(resp.Body).Decode(&result) // Decodifica a resposta JSON
	if err != nil {
		return PlayerPosition{}, err
	}

	return result.Player, nil
}
