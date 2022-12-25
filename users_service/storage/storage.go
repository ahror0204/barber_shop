package storage

import (
	"github.com/barber_shop/users_service/storage/postgres"
	"github.com/barber_shop/users_service/storage/repo"
	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Customer() repo.CustomerStorageI
	Salon() repo.SalonStorageI
	Staff() repo.StaffStorageI
	Permission() repo.PermissionStorageI
}

type storagePg struct {
	customerRepo repo.CustomerStorageI
	salonRepo    repo.SalonStorageI
	staffRepo    repo.StaffStorageI
	permissionRepo repo.PermissionStorageI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &storagePg{
		customerRepo: postgres.NewCustomerRepo(db),
		salonRepo:    postgres.NewSalonRepo(db),
		staffRepo:    postgres.NewStaffRepo(db),
		permissionRepo: postgres.NewPermission(db),
	}
}

func (s *storagePg) Customer() repo.CustomerStorageI {
	return s.customerRepo
}

func (s *storagePg) Salon() repo.SalonStorageI {
	return s.salonRepo
}

func (s *storagePg) Staff() repo.StaffStorageI {
	return s.staffRepo
}

func (s *storagePg) Permission() repo.PermissionStorageI {
	return s.permissionRepo
}