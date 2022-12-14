package service

import (
	"context"

	pbu "github.com/barber_shop/users_service/genproto/users_service"
	l "github.com/barber_shop/users_service/pkg/logger"
	"github.com/barber_shop/users_service/storage"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type SalonService struct {
	pbu.UnimplementedSalonServiceServer
	storage storage.StorageI
	logger  l.Logger
}

func NewSalonService(db *sqlx.DB, log l.Logger) *SalonService {
	return &SalonService{
		storage: storage.NewStoragePg(db),
		logger:  log,
	}
}

func (s *SalonService) CreateSalon(ctx context.Context, req *pbu.Salon) (*pbu.Salon, error) {
	salon, err := s.storage.Salon().CreateSalon(req)
	if err != nil {
		s.logger.Error("failed while creating salon", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating salon")
	}

	return salon, nil
}

func (s *SalonService) UpdateSalon(ctx context.Context, req *pbu.Salon) (*pbu.Salon, error) {
	salon, err := s.storage.Salon().UpdateSalon(req)
	if err != nil {
		s.logger.Error("failed while updating salon", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while updating salon")
	}
	return salon, nil
}

func (s *SalonService) GetSalonByID(ctx context.Context, req *pbu.ID) (*pbu.Salon, error) {
	salon, err := s.storage.Salon().GetSalonByID(req)
	if err != nil {
		s.logger.Error("failed while getting salon by id", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting salon by id")
	}
	return salon, nil
}

func (s *SalonService) GetListSalons(ctx context.Context, req *pbu.GetListParams) (*pbu.AllSalons, error) {
	salons, err := s.storage.Salon().GetListSalons(req)
	if err != nil {
		s.logger.Error("failed while gting salons list", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while failed while gting salons list")
	}
	return salons, nil
}

func (s *SalonService) DeleteSalon(ctx context.Context, req *pbu.ID) (*pbu.Empty, error) {
	err := s.storage.Salon().DeleteSalon(req)
	if err != nil {
		s.logger.Error("failed while deleting salon", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while deleting salon")
	}
	return &pbu.Empty{}, nil
}
