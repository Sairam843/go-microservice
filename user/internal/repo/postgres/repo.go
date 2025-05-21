package postgres

import (
	"github.com/jackc/pgx/v4/pgxpool"
	"go.uber.org/zap"
)

type Repo struct {
	logger *zap.Logger
	client *pgxpool.Pool
}

func NewRepo(pgxpool *pgxpool.Pool, zapLogger *zap.Logger) *Repo {
	return &Repo{
		client: pgxpool,
		logger: zapLogger,
	}
}
