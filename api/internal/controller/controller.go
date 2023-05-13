package controller

import (
	"github.com/atlant1da-404/url_shorter/config"
	"github.com/atlant1da-404/url_shorter/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Options is used to parameterize http controller via NewHTTPController.
type Options struct {
	Handler  *gin.Engine
	Services service.Services
	Config   *config.Config
}

// RouterContext provides a shared context for all routers.
type RouterContext struct {
	service service.Services
	cfg     *config.Config
}

// RouterOptions provides shared options for all routers.
type RouterOptions struct {
	Handler  *gin.RouterGroup
	Services service.Services
	Config   *config.Config
}

// New is used to create new http controllers.
func New(options Options) {
	options.Handler.Use(
		corsMiddleware,
	)

	routerOptions := RouterOptions{
		Handler:  options.Handler.Group("/"),
		Services: options.Services,
		Config:   options.Config,
	}

	// routers
	{
		setupURLRoutes(routerOptions)
	}
}

// corsMiddleware - used to allow incoming cross-origin requests.
func corsMiddleware(c *gin.Context) {
	c.Header("Access-Control-Allow-Origin", "*")
	c.Header("Access-Control-Allow-Methods", "*")
	c.Header("Access-Control-Allow-Headers", "*")
	c.Header("Content-Type", "application/json")

	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusOK)
	}
}
