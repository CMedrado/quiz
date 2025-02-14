package cmd

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

type answerQuestionRequest struct {
	NumberQuestion int `json:"number_question" validate:"required"`
	Answer         int `json:"answer" validate:"required"`
}

type answerQuestionResponse struct {
	IsCorrect bool `json:"is_correct"`
}

func answerQuestion(ctx context.Context, username string, questionID, answerIndex int) (bool, error) {
	url := fmt.Sprintf("%s/question/?user=%s", apiBaseURL, username)

	request := answerQuestionRequest{
		NumberQuestion: questionID,
		Answer:         answerIndex,
	}

	requestBody, _ := json.Marshal(request)

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	var result answerQuestionResponse
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return false, err
	}

	return result.IsCorrect, nil
}
