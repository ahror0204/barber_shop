package service

import (
	"context"
	"database/sql"
	"errors"

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

func (c *UsersService) CreateCustomer(ctx context.Context, req *pb.Customer) (*pb.Customer, error) {
	customer, err := c.storage.Customer().CreateCustomer(req)
	if err != nil {
		c.logger.Error("failed while creating customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating customer")
	}

	return customer, nil
}

func (c *UsersService) UpdateCustomer(ctx context.Context, req *pb.Customer) (*pb.Customer, error) {
	customer, err := c.storage.Customer().UpdateCustomer(req)
	if err != nil {
		c.logger.Error("failed while updating customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating customer")
	}
	return customer, nil
}

func (c *UsersService) GetCustomerByID(ctx context.Context, req *pb.ID) (*pb.Customer, error) {
	customer, err := c.storage.Customer().GetCustomerByID(req)
	if err != nil {
		c.logger.Error("failed while getting customer by id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting customer by id")
	}
	return customer, nil
}

func (c *UsersService) GetListCustomers(ctx context.Context, req *pb.GetCustomerParams) (*pb.AllCustomers, error) {
	Customers, err := c.storage.Customer().GetListCustomers(req)
	if err != nil {
		c.logger.Error("failed while gting customers list", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while failed while gting customers list")
	}
	return Customers, nil
}

func (c *UsersService) DeleteCustomer(ctx context.Context, req *pb.ID) (*pb.Empty, error) {
	err := c.storage.Customer().DeleteCustomer(req)
	if err != nil {
		c.logger.Error("failed while deleting customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting customer")
	}
	return nil, nil
}

func (c *UsersService) GetCustomerByEmail(ctx context.Context, req *pb.Email) (*pb.Customer, error) {
	customer, err := c.storage.Customer().GetCustomerByEmail(req)
	if err != nil {
		c.logger.Error("failed while getting customer by email", l.Error(err))
		if !errors.Is(err, sql.ErrNoRows) {
			return nil, status.Error(codes.NotFound, err.Error())
		}
		return nil, status.Error(codes.Internal, "failed while getting customer by email")
	}

	return customer, nil
}

func (c *UsersService) UpdateCustomerPassword(ctx context.Context, req *pb.UpdateCustomerPasswordRequest) (*pb.Empty, error) {
	err := c.storage.Customer().UpdateCustomerPassword(req)
	if err != nil {
		c.logger.Error("failed while updating customer password", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating customer password")
	}

	return &pb.Empty{}, nil
}