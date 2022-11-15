package repo

import s "github.com/barber_shop/users/structures"

type UserStorageI interface{
	CreateUser(*s.User) (string, error)
}