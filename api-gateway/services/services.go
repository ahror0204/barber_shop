package services

import (
	"fmt"

	config "github.com/barber_shop/api-gateway/config"
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

type IServiceManager interface {
	CustomerAuthService() pbu.CustomerAuthServiceClient
	CustomerService() pbu.CustomerServiceClient
	SalonService() pbu.SalonServiceClient
	StaffService() pbu.StaffServiceClient
	StaffAuthService() pbu.StaffAuthServiceClient
}

type serviceManager struct {
	cfg         config.Config
	connections map[string]interface{}
}

func NewServiceManager(conf *config.Config) (IServiceManager, error) {
	resolver.SetDefaultScheme("dns")
	connUsersService, err := grpc.Dial(
		fmt.Sprintf("%s:%d", conf.UserServiceHost, conf.UserServicePort),
		grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	serviceManager := &serviceManager{
		cfg: *conf,
		connections: map[string]interface{}{
			"customer_auth_service": pbu.NewCustomerAuthServiceClient(connUsersService),
			"customer_service":      pbu.NewCustomerServiceClient(connUsersService),
			"salon_service":         pbu.NewSalonServiceClient(connUsersService),
			"staff_service":         pbu.NewStaffServiceClient(connUsersService),
			"staff_auth_service":    pbu.NewStaffAuthServiceClient(connUsersService),
		},
	}
	return serviceManager, nil
}

func (s *serviceManager) CustomerAuthService() pbu.CustomerAuthServiceClient {
	return s.connections["customer_auth_service"].(pbu.CustomerAuthServiceClient)
}

func (s *serviceManager) CustomerService() pbu.CustomerServiceClient {
	return s.connections["customer_service"].(pbu.CustomerServiceClient)
}

func (s *serviceManager) SalonService() pbu.SalonServiceClient {
	return s.connections["salon_service"].(pbu.SalonServiceClient)
}

func (s *serviceManager) StaffService() pbu.StaffServiceClient {
	return s.connections["staff_service"].(pbu.StaffServiceClient)
}

func (s *serviceManager) StaffAuthService() pbu.StaffAuthServiceClient {
	return s.connections["staff_auth_service"].(pbu.StaffAuthServiceClient)
}
