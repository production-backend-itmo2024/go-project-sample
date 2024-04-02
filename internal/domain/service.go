package domain

import (
	"go-sample/internal/pkg/client/weather"
	"go-sample/internal/repository"
)

type Service struct {
	weatherClient *weather.Client
	repo          *repository.Repository
}

func NewService(
	weatherClient *weather.Client,
	repo *repository.Repository,
) *Service {
	return &Service{
		weatherClient: weatherClient,
		repo:          repo,
	}
}
