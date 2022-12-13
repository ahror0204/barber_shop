package repo

import pbu "github.com/barber_shop/users_service/genproto/users_service"

type SalonStorageI interface {
	CreateSalon(*pbu.Salon) (*pbu.Salon, error)
	UpdateSalon(*pbu.Salon) (*pbu.Salon, error)
	GetSalonByID(*pbu.ID) (*pbu.Salon, error)
	GetListSalons(*pbu.GetSalonsParams) (*pbu.AllSalons, error)
	DeleteSalon(*pbu.ID) error
}
