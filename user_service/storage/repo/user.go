package repo

import pb "github.com/barber_shop/user_service/genproto"

type UserStorageI interface {
	CreateUser(*pb.User) (*pb.ID, error)
	UpdateUser(*pb.User) (*pb.User, error)
	GetUserByID(*pb.ID) (*pb.User, error)
	GetAllUsers(*pb.GetUserParams) (*pb.AllUsers, error)
	DeleteUser(*pb.ID) error
}
