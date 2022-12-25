package repo

import pbu "github.com/barber_shop/users_service/genproto/users_service"

type StaffStorageI interface {
	CreateStaff(*pbu.Staff) (*pbu.Staff, error)
	UpdateStaff(*pbu.Staff) (*pbu.Staff, error)
	GetStaffByID(*pbu.ID) (*pbu.Staff, error)
	GetListStaff(*pbu.GetListParams) (*pbu.ListStaff, error)
	DeleteStaff(*pbu.ID) error
	GetStaffByEmail(*pbu.Email) (*pbu.Staff, error)
	UpdateStaffPassword(*pbu.UpdatePasswordRequest) (error)
}
