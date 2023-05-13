package service

import (
	"context"
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/atlant1da-404/url_shorter/config"
)

type urlService struct {
	cfg      *config.Config
	storages Storages
}

func NewUrlService(storages Storages, cfg *config.Config) UrlService {
	return &urlService{storages: storages, cfg: cfg}
}

func (u *urlService) GenerateShortUrl(ctx context.Context, opt *GenerateShortURLOptions) (string, error) {
	if opt.URL == "" {
		return "", errors.New("url not found")
	}

	uniqKey, err := u.generateUniqKey(u.cfg.App.URLLen)
	if err != nil {
		return "", fmt.Errorf("failed to generate uniq key: %w", err)
	}

	err = u.storages.URLStorage.SaveURL(ctx, &SaveURLOptions{Key: uniqKey, URL: opt.URL})
	if err != nil {
		return "", fmt.Errorf("failed to save url: %w", err)
	}

	shortUrl := u.generateURL(uniqKey)
	if shortUrl == "" {
		return "", fmt.Errorf("failed to generate url")
	}

	return shortUrl, nil
}

func (u *urlService) generateUniqKey(len int) (string, error) {
	b := make([]byte, len)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%X", b), nil
}

func (u *urlService) generateURL(uniqKey string) string {
	return fmt.Sprintf("%s/%s", u.cfg.App.BaseURL, uniqKey)
}

func (u *urlService) GetURL(ctx context.Context, opt *GetURLOptions) (string, error) {
	if opt.Key == "" {
		return "", errors.New("key not found")
	}

	url, err := u.storages.URLStorage.GetURL(ctx, opt.Key)
	if err != nil {
		return "", fmt.Errorf("failed to get original url: %w", err)
	}
	if url == "" {
		return "", errors.New("original url not found")
	}

	return url, nil
}
