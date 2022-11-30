package main

import (
	"exam/api-gateway/api"
	_ "exam/api-gateway/api/docs"
	"exam/api-gateway/config"
	"exam/api-gateway/pkg/logger"
	"exam/api-gateway/services"
	"fmt"
)

func main() {
	fmt.Println("hello world 13")

	cfg := config.Load()
	fmt.Println("hello world 17")

	log := logger.New(cfg.LogLevel, "api_gateway")
	fmt.Println("hello world 17")

	serviceManager, err := services.NewServiceManager(&cfg)
	if err != nil {
		log.Error("gRPC dial error", logger.Error(err))
	}
	fmt.Println("hello world 20")

	server := api.New(api.Option{
		Conf:           cfg,
		Logger:         log,
		ServiceManager: serviceManager,
	})
	fmt.Println("hello world 26")
	if err := server.Run(cfg.HTTPPort); err != nil {
		log.Fatal("failed to run http server", logger.Error(err))
		panic(err)
	}
}
