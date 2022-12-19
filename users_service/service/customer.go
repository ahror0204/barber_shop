package service

import (
	"context"
	"database/sql"
	"errors"

	pbu "github.com/barber_shop/users_service/genproto/users_service"
	l "github.com/barber_shop/users_service/pkg/logger"
	"github.com/barber_shop/users_service/storage"
	"github.com/jmoiron/sqlx"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type CustomerService struct {
	pbu.UnimplementedCustomerServiceServer
	storage storage.StorageI
	logger  l.Logger
}

func NewCustomerService(db *sqlx.DB, log l.Logger) *CustomerService {
	return &CustomerService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (c *CustomerService) CreateCustomer(ctx context.Context, req *pbu.Customer) (*pbu.Customer, error) {
	customer, err := c.storage.Customer().CreateCustomer(req)
	if err != nil {
		c.logger.Error("failed while creating customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating customer")
	}

	return customer, nil
}

func (c *CustomerService) UpdateCustomer(ctx context.Context, req *pbu.Customer) (*pbu.Customer, error) {
	customer, err := c.storage.Customer().UpdateCustomer(req)
	if err != nil {
		c.logger.Error("failed while updating customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating customer")
	}
	return customer, nil
}

func (c *CustomerService) GetCustomerByID(ctx context.Context, req *pbu.ID) (*pbu.Customer, error) {
	customer, err := c.storage.Customer().GetCustomerByID(req)
	if err != nil {
		c.logger.Error("failed while getting customer by id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting customer by id")
	}
	return customer, nil
}

func (c *CustomerService) GetListCustomers(ctx context.Context, req *pbu.GetListParams) (*pbu.AllCustomers, error) {
	Customers, err := c.storage.Customer().GetListCustomers(req)
	if err != nil {
		c.logger.Error("failed while gting customers list", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while failed while gting customers list")
	}
	return Customers, nil
}

func (c *CustomerService) DeleteCustomer(ctx context.Context, req *pbu.ID) (*pbu.Empty, error) {
	err := c.storage.Customer().DeleteCustomer(req)
	if err != nil {
		c.logger.Error("failed while deleting customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting customer")
	}
	return &pbu.Empty{}, nil
}

func (c *CustomerService) GetCustomerByEmail(ctx context.Context, req *pbu.Email) (*pbu.Customer, error) {
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

func (c *CustomerService) UpdateCustomerPassword(ctx context.Context, req *pbu.UpdatePasswordRequest) (*pbu.Empty, error) {
	err := c.storage.Customer().UpdateCustomerPassword(req)
	if err != nil {
		c.logger.Error("failed while updating customer password", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating customer password")
	}

	return &pbu.Empty{}, nil
}
