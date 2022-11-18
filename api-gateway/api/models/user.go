package models

import (
	"time"

	pb "github.com/barber_shop/api-gateway/genproto"
)

type User struct {
	ID          string    `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"passward"`
	Gender      string    `json:"gender"`
	ImageURL    string    `json:"image_url"`
	CreatedAT   time.Time `json:"created_at"`
	UpdatedAT   time.Time `json:"updated_at"`
	DeletedAT   time.Time `json:"deleted_at"`
}

type CreateUserRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	Password    string `json:"passward"`
	Gender      string `json:"gender"`
	ImageURL    string `json:"image_url"`
}

type GetAllParams struct {
	Limit  int    `json:"limit" binding:"required" default:"10"`
	Page   int    `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}

type CreateUserRespons struct{
	ID string `json:"id"`
}

func ParsUserStruct(user User) pb.User {
	return pb.User{
		FirstName: user.FirstName,
		LastName: user.LastName,
		PhoneNumber: user.PhoneNumber,
		Email: user.Email,
		UserName: user.UserName,
		Password: user.Password,
		Gender: user.Gender,
		ImageUrl: user.ImageURL,
	}
}
