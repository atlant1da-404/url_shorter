package service

import "context"

// Storages contains all available storages
type Storages struct {
	URLStorage URLStorage
}

type URLStorage interface {
	SaveURL(ctx context.Context, opt *SaveURLOptions) error
	GetURL(ctx context.Context, key string) (string, error)
}

type SaveURLOptions struct {
	Key string
	URL string
}
