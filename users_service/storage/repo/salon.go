package repo

import pb "github.com/barber_shop/users_service/genproto"

type SalonStorageI interface {
	CreateSalon(*pb.Salon) (*pb.ID, error)
	UpdateSalon(*pb.Salon) (*pb.Salon, error)
	GetSalonByID(*pb.ID) (*pb.Salon, error)
	GetListSalons(*pb.GetSalonsParams) (*pb.AllSalons, error)
	DeleteSalon(*pb.ID) error
}
