package repo

import s "github.com/barber_shop/users/structures"

type UserStorageI interface{
	CreateUser(*s.User) (string, error)
	UpdateUser(*s.User) (*s.User, error)
	GetUserByID(string) (*s.User, error)
	GetAllUsers(*s.GetUsersParams)(*s.AllUsers, error)
	DeleteUser(id string) error

}