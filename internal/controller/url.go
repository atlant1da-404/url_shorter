package controller

import (
	"github.com/atlant1da-404/url_shorter/internal/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type urlRouter struct {
	RouterContext
}

// setupURLRoutes implements new URL handler.
func setupURLRoutes(options RouterOptions) {
	r := &urlRouter{
		RouterContext{
			service: options.Services,
			cfg:     options.Config,
		},
	}

	options.Handler.POST("/short", r.shortURL)
	options.Handler.GET("/:key", r.redirectURL)
}

type shortURLRequestBody struct {
	URL string `json:"url"`
}

type shortURLResponseBody struct {
	URL string `json:"url"`
}

func (u *urlRouter) shortURL(c *gin.Context) {
	var body shortURLRequestBody
	err := c.ShouldBindJSON(&body)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	shortURL, err := u.service.UrlService.GenerateShortUrl(c, &service.GenerateShortURLOptions{URL: body.URL})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, &shortURLResponseBody{shortURL})
}

type redirectURLRequestBody struct {
	Key string `uri:"key"`
}

func (u *urlRouter) redirectURL(c *gin.Context) {
	var params redirectURLRequestBody
	err := c.ShouldBindUri(&params)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, err.Error())
		return
	}

	originalURL, err := u.service.UrlService.GetURL(c, &service.GetURLOptions{Key: params.Key})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.Redirect(http.StatusPermanentRedirect, originalURL)
}
