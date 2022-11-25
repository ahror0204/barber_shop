package main

import (
	"net"

	c "github.com/barber_shop/users_service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/barber_shop/users_service/config"
	pb "github.com/barber_shop/users_service/genproto"
	"github.com/barber_shop/users_service/pkg/db"
	logger "github.com/barber_shop/users_service/pkg/logger"
)

func main() {
	cfg := config.Load("./")

	log := logger.New(cfg.LogLevel, "users-service")
	defer logger.Cleanup(log)
	
	db, err := db.ConnectToDB(cfg)
	if err != nil {
        log.Fatal("sqlx connection to postgres error", logger.Error(err))
    }

	userService := c.NewUsersService(db, log)

	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterCustomerServiceServer(s, userService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
