package main

import (
	"net"

	c "github.com/barber_shop/user_service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/barber_shop/user_service/config"
	pb "github.com/barber_shop/user_service/genproto"
	logger "github.com/barber_shop/user_service/pkg/logger"
	"github.com/barber_shop/user_service/pkg/db"	
)

func main() {
	cfg := config.Load("./")

	log := logger.New(cfg.LogLevel, "template-service")
	defer logger.Cleanup(log)
	
	db, err := db.ConnectToDB(cfg)
	if err != nil {
        log.Fatal("sqlx connection to postgres error", logger.Error(err))
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
