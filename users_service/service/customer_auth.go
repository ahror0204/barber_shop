package service

import (
	"github.com/barber_shop/users_service/config"
	pbu "github.com/barber_shop/users_service/genproto/users_service"
	"github.com/barber_shop/users_service/pkg/logger"
	"github.com/barber_shop/users_service/storage"
)

type CustomerAuthService struct {
	pbu.UnimplementedCustomerAuthServiceServer
	storage storage.StorageI
	inMemory storage.InMemoryStorageI
	cfg      *config.Config
	logger   *logger.Logger
}

func NewCustomerAuthService(s storage.StorageI, inMemory storage.InMemoryStorageI, cfg *config.Config, logger *logger.Logger) *CustomerAuthService {
	return &CustomerAuthService{
		storage:  s,
		inMemory: inMemory,
		cfg:      cfg,
		logger:   logger,
	}
}

const (
	RegisterCodeKey   = "register_code_"
	ForgotPasswordKey = "forgot_password_code_"
)
