package main

import (
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"

	"quiz/app/http"
	"quiz/app/http/question"
	rankHandler "quiz/app/http/rank"
	"quiz/config"
	"quiz/domain/usecases/questions"
	"quiz/domain/usecases/rank"
	redisGateway "quiz/infrastructure/gateways/redis"
	repository "quiz/infrastructure/memory"
	"quiz/pkg"
)

func main() {
	logger := pkg.CreateLogger()

	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Fatal("unable to load app configuration", zap.Error(err))
	}
	logger = logger.With(zap.String("app_name", cfg.PodName))

	client := redis.NewClient(&redis.Options{
		Addr:     cfg.RedisAddr,
		Password: cfg.RedisPassword,
		DB:       cfg.RedisDB,
	})

	questionStorage := repository.NewMemoryQuestionStorage()
	cacheClient := redisGateway.NewRedisGateway(client, cfg.RedisKey)

	getQuestionUseCase := questions.NewGetQuestionUseCase(questionStorage, cacheClient)
	submitAnswerUseCase := questions.NewAnswerQuestionUseCase(cacheClient, questionStorage)
	questionHandler := question.NewQuestionHandler(getQuestionUseCase, submitAnswerUseCase)

	getRankUseCase := rank.NewGetLeaderboardUseCase(cacheClient)
	getPlayerPositionUseCase := rank.NewGetPlayerPositionUseCase(cacheClient)
	rankHandler := rankHandler.NewRankHandler(getRankUseCase, getPlayerPositionUseCase)

	api := http.NewAPI(rankHandler, questionHandler, logger)
	logger.Info("Starting API", zap.String("host", "0.0.0.0"), zap.Int("port", cfg.Port))
	api.Start("0.0.0.0", cfg.Port)
}
