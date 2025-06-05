package main

import (
	"log"
	"net"
	"os"

	userGRPCHandler "movie/user/api/grpc"
	postgresDB "movie/user/internal/db/postgres"
	userPB "movie/user/internal/pb/user"
	postgresRepo "movie/user/internal/repo/postgres"
	"movie/user/internal/service"

	"github.com/joho/godotenv"
	"go.uber.org/zap"
	"google.golang.org/grpc"
)

const (
	FALLBACK_PORT = "8090"
)

func main() {
	var (
		logger *zap.Logger
	)

	err := godotenv.Load()
	if err != nil {
		log.Fatalf("unable to load .env, %v", err)
	}

	if os.Getenv("APP_ENV") == "PROD" {
		logger, err = zap.NewProduction()
	} else {
		logger, err = zap.NewDevelopment()
	}

	if err != nil {
		log.Fatalf("Failed to init logger: %v", err)
	}

	defer logger.Sync()

	// Init DB client
	dbClient := postgresDB.NewPostgresClient(logger)

	// Init repo
	repo := postgresRepo.NewRepo(dbClient, logger)

	// Init service
	svc := service.NewService(repo, logger)

	// Init handler with service
	handler := userGRPCHandler.NewUserGRPCHandler(svc)

	// Start gRPC server
	port := os.Getenv("APP_PORT")
	if port == "" {
		port = FALLBACK_PORT // default fallback
	}

	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logger.Fatal("Failed to listen", zap.Error(err))
	}

	grpcServer := grpc.NewServer()
	userPB.RegisterUserServer(grpcServer, handler)

	logger.Info("gRPC server started", zap.String("port", port))
	if err := grpcServer.Serve(lis); err != nil {
		logger.Fatal("Failed to serve", zap.Error(err))
	}
}
