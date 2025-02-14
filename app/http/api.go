package http

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"quiz/app/http/question"
	"quiz/app/http/rank"
	"quiz/pkg"
)

// API struct to manage ranking and questions API.
type API struct {
	rank     *rank.Handler
	question *question.Handler
	log      *zap.Logger
}

// NewAPI initializes a new API instance.
func NewAPI(rank *rank.Handler, question *question.Handler, logger *zap.Logger) *API {
	return &API{
		rank:     rank,
		question: question,
		log:      logger,
	}
}

// Start initializes and starts the API server.
func (a *API) Start(host string, port int) {
	r := chi.NewRouter()

	r.Use(requestLogger(a.log))

	// Rank endpoints
	r.Route("/rank", func(r chi.Router) {
		r.Get("/player/", a.rank.GetPlayerPosition)
		r.Get("/", a.rank.GetLeaderboard)
	})

	// Question endpoints
	r.Route("/question", func(r chi.Router) {
		r.Get("/", a.question.GetQuestion)
		r.Post("/", a.question.AnswerQuestion)
	})

	address := fmt.Sprintf("%s:%d", host, port)

	server := &http.Server{
		Handler:      r,
		Addr:         address,
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  30 * time.Second,
	}

	a.log.Info("Starting API", zap.String("address", address))
	if err := server.ListenAndServe(); err != nil {
		a.log.Fatal("Failed to start server", zap.Error(err))
	}
}

// requestLogger adds logging middleware.
func requestLogger(log *zap.Logger) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			r = r.WithContext(pkg.WithCtx(r.Context(), log))
			next.ServeHTTP(w, r)
		})
	}
}
