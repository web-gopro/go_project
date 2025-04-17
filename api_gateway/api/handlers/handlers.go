package handlers

import (
	"github.com/saidamir98/udevs_pkg/logger"
	"gitlab.com/e-market724/back-end/api_gateway/redis"
	"gitlab.com/e-market724/back-end/api_gateway/services"
)

type Handlers struct {
	log      logger.LoggerI
	services services.ServiceManagerI
	redis    redis.RedisRepoI
}

func NewHandler(log logger.LoggerI, services services.ServiceManagerI, redis redis.RedisRepoI) Handlers {

	return Handlers{log: log, services: services, redis: redis}
}
