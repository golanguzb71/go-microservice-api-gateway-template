package main

import (
	rediscache "github.com/golanguzb70/redis-cache"
	"gitlab.udevs.io/eld/eld_go_api_gateway/api"
	"gitlab.udevs.io/eld/eld_go_api_gateway/config"
	"gitlab.udevs.io/eld/eld_go_api_gateway/pkg/logger"
	"gitlab.udevs.io/eld/eld_go_api_gateway/services"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "livesphere_go_api_gateway")
	gprcClients, err := services.NewGrpcClients(&cfg)
	if err != nil {
		log.Error("Error while getting grpc clients", logger.Error(err))
		return
	}

	cache := rediscache.New(&rediscache.Config{
		RedisHost: cfg.RedisHost,
		RedisPort: cfg.RedisPort,
	})

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      &cfg,
		Services: gprcClients,
		Cache:    cache,
	})

	server.Run(":" + cfg.HttpPort)
}
