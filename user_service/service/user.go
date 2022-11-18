package service

import (
	"context"
	pb "github.com/barber_shop/user_service/genproto"
	"github.com/barber_shop/user_service/storage"
	"github.com/jmoiron/sqlx"
)


type UserService struct {
	storage storage.StorageI
}

func NewUserService(db *sqlx.DB) *UserService{
	return &UserService{
		storage: storage.NewStoragePg(db),
	}
}

func (u *UserService) CreateUser(ctx context.Context, req *pb.User) (*pb.ID, error){
	id, err := u.storage.User().CreateUser(req)
	if err != nil {
		return nil, err
	}

	return id, nil
}

func (u *UserService) UpdateUser(ctx context.Context, req *pb.User) (*pb.User, error){
	_, _ = u.storage.User().UpdateUser(req)
	return nil, nil
}

func (u *UserService) GetUserByID(ctx context.Context, req *pb.ID) (*pb.User, error){
	_, _ = u.storage.User().GetUserByID(req)
	return nil, nil
}

func (u *UserService) GetAllUsers(ctx context.Context, req *pb.GetUserParams) (*pb.AllUsers, error){
	_, _ = u.storage.User().GetAllUsers(req)
	return nil, nil
}

func (u *UserService) DeleteUser(ctx context.Context, req *pb.ID) (*pb.Empty, error) {
	err := u.storage.User().DeleteUser(req)
	if err != nil {
		return nil, err
	}
	return nil, nil
}
