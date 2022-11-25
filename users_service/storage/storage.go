package storage

import (
	"github.com/barber_shop/users_service/storage/postgres"
	"github.com/barber_shop/users_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Customer() repo.CustomerStorageI
	Salon() repo.SalonStorageI
}

type storagePg struct {
	customerRepo repo.CustomerStorageI
	salonRepo    repo.SalonStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		customerRepo: postgres.NewCustomer(db),
	}
}

func (s *storagePg) Customer() repo.CustomerStorageI {
	return s.customerRepo
}

func (s *storagePg) Salon() repo.SalonStorageI {
	return s.salonRepo
}
