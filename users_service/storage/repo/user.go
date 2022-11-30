package repo

import pb "github.com/barber_shop/users_service/genproto"

type CustomerStorageI interface {
	CreateCustomer(*pb.Customer) (*pb.ID, error)
	UpdateCustomer(*pb.Customer) (*pb.Customer, error)
	GetCustomerByID(*pb.ID) (*pb.Customer, error)
	GetListCustomers(*pb.GetCustomerParams) (*pb.AllCustomers, error)
	DeleteCustomer(*pb.ID) error
	GetCustomerByEmail(*pb.Email) (*pb.Customer, error)
}
