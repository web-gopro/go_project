package main

import (
	"context"
	"log"

	"github.com/saidamir98/udevs_pkg/logger"
	"gitlab.com/e-market724/back-end/api_gateway/api"
	"gitlab.com/e-market724/back-end/api_gateway/config"
	"gitlab.com/e-market724/back-end/api_gateway/pkg/db"
	"gitlab.com/e-market724/back-end/api_gateway/redis"
	"gitlab.com/e-market724/back-end/api_gateway/services"
)

var ctx = context.Background()

func main() {

	logr := logger.NewLogger("", logger.LevelDebug)
	cfg := config.Load()

	service := services.Service()

	log.Println(service)

	rediCli, err := db.ConnRedis(logr, context.Background(), cfg.RedisConfig)

	if err != nil {

		logr.Debug(err.Error())
		return

	}

	cache := redis.NewRedisRepo(rediCli, logr)

	engine := api.Api(api.Options{Services: service, Redis: cache, Log: logr})

	engine.Run(":8080")

}
