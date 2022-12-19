package repo

import pbu "github.com/barber_shop/users_service/genproto/users_service"

type CustomerStorageI interface {
	CreateCustomer(*pbu.Customer) (*pbu.Customer, error)
	UpdateCustomer(*pbu.Customer) (*pbu.Customer, error)
	GetCustomerByID(*pbu.ID) (*pbu.Customer, error)
	GetListCustomers(*pbu.GetListParams) (*pbu.AllCustomers, error)
	DeleteCustomer(*pbu.ID) error
	GetCustomerByEmail(*pbu.Email) (*pbu.Customer, error)
	UpdateCustomerPassword(*pbu.UpdatePasswordRequest) (error)
}
