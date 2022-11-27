package services

import (
	"fmt"

	config "github.com/barber_shop/api-gateway/config"
	pb "github.com/barber_shop/api-gateway/genproto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	UserService() pb.CustomerServiceClient
}

type serviceManager struct {
	userService pb.CustomerServiceClient
}

func (s *serviceManager) UserService() pb.CustomerServiceClient {
	return s.userService
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")
	connUser, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}
	serviceManager := &serviceManager{
		userService: pb.NewCustomerServiceClient(connUser),
	}

	return serviceManager, nil
}
