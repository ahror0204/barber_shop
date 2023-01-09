package models

import (
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
)

type VerifyRequest struct {
	Email string `json:"email"`
	Code  string `json:"code"`
}

type LogInRequest struct {
	Email    string `json:"email"`
	Password string `json:"passward"`
}

type ForgotPasswordRequest struct {
	Email string `json:"email"`
}

type UpdatePasswordRequest struct {
	Password string `json:"password" binding:"required"`
}

type CustomerLogInResponse struct {
	ID        string `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	Type      string `json:"type"`
	Token     string `json:"token"`
}

type AuthResponse struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	Password    string `json:"passward"`
	Gender      string `json:"gender"`
	Type        string `json:"type"`
	ImageURL    string `json:"image_url"`
	CreatedAT   string `json:"created_at"`
	Token       string `json:"token"`
}

func ParsAuthResponseToPbCustomer(c *pbu.Customer) *AuthResponse {
	return &AuthResponse{
		ID:          c.Id,
		FirstName:   c.FirstName,
		LastName:    c.LastName,
		PhoneNumber: c.PhoneNumber,
		Email:       c.Email,
		UserName:    c.UserName,
		Password:    c.Password,
		Gender:      c.Gender,
		Type:        c.Type,
		ImageURL:    c.ImageUrl,
		CreatedAT:   c.CreatedAt,
	}
}
