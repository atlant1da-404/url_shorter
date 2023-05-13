package main

import (
	"github.com/atlant1da-404/url_shorter/config"
	"github.com/atlant1da-404/url_shorter/internal/app"
)

func main() {
	cfg := config.Get()

	app.Run(cfg)
}
