package main

import (
	"fmt"
	"net"

	"github.com/Asliddin3/review-servise/config"
	pb "github.com/Asliddin3/review-servise/genproto/review"
	"github.com/Asliddin3/review-servise/pkg/db"
	"github.com/Asliddin3/review-servise/pkg/logger"
	"github.com/Asliddin3/review-servise/service"
	grpcClient "github.com/Asliddin3/review-servise/service/grpc_client"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	cfg := config.Load()

	log := logger.New(cfg.LogLevel, "")
	defer logger.Cleanup(log)

	log.Info("main:sqlxConfig",
		logger.String("host", cfg.PostgresHost),
		logger.Int("port", cfg.PostgresPort),
		logger.String("database", cfg.PostgresDatabase))
	connDb, err := db.ConnectToDb(cfg)
	fmt.Println("error inr review", err)
	if err != nil {
		log.Fatal("sqlx connection to postgres error", logger.Error(err))
	}
	grpcClient, err := grpcClient.New(cfg)
	if err != nil {
		log.Fatal("error while connect to clients", logger.Error(err))
	}

	postService := service.NewReviewService(grpcClient, connDb, log)
	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
	s := grpc.NewServer()
	reflection.Register(s)
	pb.RegisterReviewServiceServer(s, postService)
	log.Info("main: server runing",
		logger.String("port", cfg.RPCPort))
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
