package service

import (
	postgresRepo "movie/user/internal/repo/postgres"

	"go.uber.org/zap"
)

type Service struct {
	repo   *postgresRepo.Repo
	logger *zap.Logger
}

func NewService(repo *postgresRepo.Repo, zapLogger *zap.Logger) *Service {
	return &Service{
		repo:   repo,
		logger: zapLogger,
	}
}
