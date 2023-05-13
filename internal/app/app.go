package app

import (
	"github.com/atlant1da-404/url_shorter/config"
	"github.com/atlant1da-404/url_shorter/internal/controller"
	"github.com/atlant1da-404/url_shorter/internal/service"
	"github.com/atlant1da-404/url_shorter/internal/storage"
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"log"
)

func Run(cfg *config.Config) {
	handler := gin.New()

	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Address,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	storages := service.Storages{
		URLStorage: storage.NewUrlStorage(rdb),
	}

	services := service.Services{
		UrlService: service.NewUrlService(storages, cfg),
	}

	controller.New(controller.Options{
		Handler:  handler,
		Services: services,
		Config:   cfg,
	})

	err := handler.Run(cfg.App.BaseURL)
	if err != nil {
		log.Fatalln(err.Error())
	}
}
