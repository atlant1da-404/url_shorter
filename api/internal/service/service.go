package service

import "context"

// Services provides a collection of business logic services.
type Services struct {
	UrlService UrlService
}

type UrlService interface {
	GenerateShortUrl(ctx context.Context, g *GenerateShortURLOptions) (string, error)
	GetURL(ctx context.Context, opt *GetURLOptions) (string, error)
}

type GenerateShortURLOptions struct {
	URL string
}

type GetURLOptions struct {
	Key string
}
