package main

import (
	"fmt"
	"net"

	c "github.com/barber_shop/user_service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	logger "github.com/barber_shop/user_service/pkg/logger"
	"github.com/barber_shop/user_service/config"
	pb "github.com/barber_shop/user_service/genproto"
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

func main() {
	cfg := config.Load("./")

	log := logger.New(cfg.LogLevel, "template-service")
    defer logger.Cleanup(log)

	connStr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s",
		cfg.Postgres.Host,
		cfg.Postgres.Port,
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.DataBase,
	)
	db, err := sqlx.Connect("postgres", connStr)
	if err != nil {
		log.Fatal("failed to open connection:", logger.Error(err))
	}

	userService := c.NewUserService(db)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterUserServiceServer(s, userService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
