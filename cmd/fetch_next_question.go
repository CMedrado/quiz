package cmd

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type GetQuestionResponse struct {
	NumberQuestion int      `json:"number_question"`
	Question       string   `json:"question"`
	Answers        []string `json:"answers"`
}

func fetchNextQuestion(ctx context.Context, username string) (GetQuestionResponse, error) {
	url := fmt.Sprintf("%s/question/?user=%s", apiBaseURL, username)
	resp, err := http.Get(url)
	if err != nil {
		return GetQuestionResponse{}, err
	}
	defer resp.Body.Close()

	var question GetQuestionResponse
	err = json.NewDecoder(resp.Body).Decode(&question)
	if err != nil {
		return GetQuestionResponse{}, err
	}

	return question, nil
}
