package main

import (
	"context"
	"fmt"

	"github.com/Sonka-bot-for-deep-sleep/common/pkg/logger"
	"github.com/Sonka-bot-for-deep-sleep/time_service/application/config"
	"github.com/Sonka-bot-for-deep-sleep/time_service/internal/infrastructure/db/postgres"
	"github.com/Sonka-bot-for-deep-sleep/time_service/internal/infrastructure/db/redis"
	"go.uber.org/zap"
)

func main() {
	ctx := context.Background()
	log, err := logger.New()
	if err != nil {
		fmt.Println(fmt.Errorf("Error create logger instance: %w", err))
		return
	}

	cfg, err := config.MustLoad()
	if err != nil {
		log.Error("Failed load config", zap.Error(err))
		return
	}

	pg, err := postgres.NewWithConn(cfg.DSN)
	defer pg.CloseConn(ctx)
	if err != nil {
		log.Error("Failed create conn to postgresql", zap.Error(err))
		return
	}

	rdb, err := redis.NewWithConn(cfg.REDIS_URL)
	defer rdb.CloseConn()
	if err != nil {
		log.Error("Failed conn to redis", zap.Error(err))
		return
	}

}
