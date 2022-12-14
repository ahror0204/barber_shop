package main

import (
	"net"

	c "github.com/barber_shop/users_service/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"github.com/barber_shop/users_service/config"
	pb "github.com/barber_shop/users_service/genproto/users_service"
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

	customerService := c.NewCustomerService(db, log)
	salonService := c.NewSalonService(db, log)
	staffService := c.NewStaffService(db, log)


	lis, err := net.Listen("tcp", cfg.RPCPort)
	if err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}

	s := grpc.NewServer()
	pb.RegisterCustomerServiceServer(s, customerService)
	pb.RegisterSalonServiceServer(s, salonService)
	pb.RegisterStaffServiceServer(s, staffService)

	log.Info("main: server running",
		logger.String("port", cfg.RPCPort))
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("Error while listening: %v", logger.Error(err))
	}
}
