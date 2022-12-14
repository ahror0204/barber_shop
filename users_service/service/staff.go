package service

import (
	"context"

	pbu "github.com/barber_shop/users_service/genproto/users_service"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	l "github.com/barber_shop/users_service/pkg/logger"
	"github.com/barber_shop/users_service/storage"
	"github.com/jmoiron/sqlx"
)

type StaffService struct {
	pbu.UnimplementedStaffServiceServer
	storage storage.StorageI
	logger  l.Logger
}

func NewStaffService(db *sqlx.DB, log l.Logger) *StaffService {
	return &StaffService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (c *StaffService) CreateStaff(ctx context.Context, req *pbu.Staff) (*pbu.Staff, error) {
	staff, err := c.storage.Staff().CreateStaff(req)
	if err != nil {
		c.logger.Error("failed while creating staff", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating staff")
	}

	return staff, nil
}

func (c *StaffService) UpdateStaff(ctx context.Context, req *pbu.Staff) (*pbu.Staff, error) {
	staff, err := c.storage.Staff().UpdateStaff(req)
	if err != nil {
		c.logger.Error("failed while updating staff", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating staff")
	}
	return staff, nil
}

func (c *StaffService) GetStaffByID(ctx context.Context, req *pbu.ID) (*pbu.Staff, error) {
	staff, err := c.storage.Staff().GetStaffByID(req)
	if err != nil {
		c.logger.Error("failed while getting staff by id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting staff by id")
	}
	return staff, nil
}

func (c *StaffService) GetListStaffs(ctx context.Context, req *pbu.GetListParams) (*pbu.ListStaff, error) {
	staff, err := c.storage.Staff().GetListStaff(req)
	if err != nil {
		c.logger.Error("failed while gting staff list", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while failed while gting staff list")
	}
	return staff, nil
}

func (c *StaffService) DeleteStaff(ctx context.Context, req *pbu.ID) (*pbu.Empty, error) {
	err := c.storage.Staff().DeleteStaff(req)
	if err != nil {
		c.logger.Error("failed while deleting staff", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting staff")
	}
	return nil, nil
}