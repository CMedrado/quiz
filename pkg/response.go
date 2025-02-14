package pkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FullError struct {
	Type  string `json:"type"`
	Title string `json:"title"`
}

func (f FullError) Error() string {
	return fmt.Sprintf("%+s; %+s", f.Type, f.Title)
}

func (f FullError) WithTitle(title string) FullError {
	return FullError{Type: f.Type, Title: title}
}

var (
	ErrInavlidUserID       = FullError{Type: "srn:error:invalid_user_id", Title: "Invalid User ID"}
	ErrInavlidAnswer       = FullError{Type: "srn:error:invalid_answer", Title: "Invalid Answer"}
	ErrInternalServerError = FullError{Type: "srn:error:internal_server_error", Title: "Internal Server Error"}
	ErrBadRequest          = FullError{Type: "srn:error:bad_request", Title: "Bad Request"}
)

func Send(w http.ResponseWriter, response interface{}, statusCode int) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	return json.NewEncoder(w).Encode(response)
}
