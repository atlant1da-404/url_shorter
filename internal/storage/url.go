package storage

import (
	"context"
	"fmt"
	"github.com/atlant1da-404/url_shorter/internal/service"
	"github.com/redis/go-redis/v9"
)

type urlStorage struct {
	rdb *redis.Client
}

func NewUrlStorage(rdb *redis.Client) service.URLStorage {
	return &urlStorage{rdb: rdb}
}

func (u *urlStorage) SaveURL(ctx context.Context, opt *service.SaveURLOptions) error {
	err := u.rdb.Set(ctx, opt.Key, opt.URL, 0).Err()
	if err != nil {
		return fmt.Errorf("failed to save url: %w", err)
	}

	return nil
}

func (u *urlStorage) GetURL(ctx context.Context, key string) (string, error) {
	url, err := u.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", fmt.Errorf("failed to get url: %w", err)
	}

	return url, nil
}
