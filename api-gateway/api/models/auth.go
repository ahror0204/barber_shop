package models

import (
	pb "github.com/barber_shop/api-gateway/genproto"
)

type VerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type LogInCustomerRequest struct {
	Email    string `json:"email"`
	Password string `json:"passward"`
}

type AuthResponse struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	Password    string `json:"passward"`
	Gender      string `json:"gender"`
	ImageURL    string `json:"image_url"`
	CreatedAT   string `json:"created_at"`
	Token       string `json:"token"`
}

func ParsAuthResponseToPbCustomer(c pb.Customer) AuthResponse {
	return AuthResponse{
		FirstName: c.FirstName, 
		LastName: c.LastName,
		PhoneNumber: c.PhoneNumber,
		Email: c.Email,
		UserName: c.UserName,
		Password: c.Password,
		Gender: c.Gender,
		ImageURL: c.ImageUrl,
		CreatedAT: c.CreatedAt,
	}
}