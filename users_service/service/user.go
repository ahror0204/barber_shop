package service

import (
	"context"

	pb "github.com/barber_shop/users_service/genproto"
	l "github.com/barber_shop/users_service/pkg/logger"
	"github.com/barber_shop/users_service/storage"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UsersService struct {
	storage storage.StorageI
	logger  l.Logger
}

func NewUsersService(db *sqlx.DB, log l.Logger) *UsersService {
	return &UsersService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (c *UsersService) CreateCustomer(ctx context.Context, req *pb.Customer) (*pb.ID, error) {
	id, err := c.storage.Customer().CreateCustomer(req)
	if err != nil {
		c.logger.Error("failed while creating customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating customer")
	}

	return id, nil
}

func (c *UsersService) UpdateCustomer(ctx context.Context, req *pb.Customer) (*pb.Customer, error) {
	_, _ = c.storage.Customer().UpdateCustomer(req)
	return nil, nil
}

func (c *UsersService) GetCustomerByID(ctx context.Context, req *pb.ID) (*pb.Customer, error) {
	_, _ = c.storage.Customer().GetCustomerByID(req)
	return nil, nil
}

func (c *UsersService) GetListCustomers(ctx context.Context, req *pb.GetCustomerParams) (*pb.AllCustomers, error) {
	_, _ = c.storage.Customer().GetListCustomers(req)
	return nil, nil
}

func (c *UsersService) DeleteCustomer(ctx context.Context, req *pb.ID) (*pb.Empty, error) {
	err := c.storage.Customer().DeleteCustomer(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
