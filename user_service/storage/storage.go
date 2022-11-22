package storage

import (
	"github.com/barber_shop/user_service/storage/postgres"
	"github.com/barber_shop/user_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	User() repo.UserStorageI
	Salon() repo.SalonStorageI
}

type storagePg struct {
	userRepo repo.UserStorageI
	salonRepo repo.SalonStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		userRepo: postgres.NewUser(db),
	}
}

func (s *storagePg) User() repo.UserStorageI {
	return s.userRepo
}

func (s *storagePg) Salon() repo.SalonStorageI {
	return s.salonRepo
}