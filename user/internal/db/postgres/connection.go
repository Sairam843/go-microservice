package postgres

import (
	"context"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/joho/godotenv"
	"go.uber.org/zap"
)

func getEnvInt(key string, defaultVal int) int {
	valStr := os.Getenv(key)
	if valStr == "" {
		return defaultVal
	}
	val, err := strconv.Atoi(valStr)
	if err != nil {
		return defaultVal
	}
	return val
}

func NewPostgresClient(logger *zap.Logger) *pgxpool.Pool {
	if err := godotenv.Load(); err != nil {
		logger.Warn("No .env file found, falling back to system env")
	}

	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	sslMode := os.Getenv("SSL_MODE")

	if dbHost == "" || dbPort == "" || dbUser == "" || dbPassword == "" || dbName == "" {
		logger.Fatal("Missing required environment variables")
	}

	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		dbUser, dbPassword, dbHost, dbPort, dbName, sslMode,
	)

	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		logger.Fatal("Unable to parse DB config", zap.Error(err))
	}

	config.MaxConns = int32(getEnvInt("DB_MAX_CONNS", 10))
	config.MaxConnLifetime = time.Duration(getEnvInt("DB_CONN_LIFETIME_SEC", 1800)) * time.Second
	config.MaxConnIdleTime = time.Duration(getEnvInt("DB_CONN_IDLE_TIME_SEC", 300)) * time.Second
	config.HealthCheckPeriod = time.Duration(getEnvInt("DB_HEALTH_CHECK_PERIOD_SEC", 60)) * time.Second

	db, err := pgxpool.ConnectConfig(context.Background(), config)
	if err != nil {
		logger.Fatal("Unable to connect to DB", zap.Error(err))
	}

	logger.Info("Connected to PostgreSQL successfully")
	return db
}
