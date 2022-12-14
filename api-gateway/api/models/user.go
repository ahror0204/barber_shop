package models

import (
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
)

type Customer struct {
	ID          string `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	Password    string `json:"password"`
	Gender      string `json:"gender"`
	Type        string `json:"type"`
	ImageURL    string `json:"image_url"`
	CreatedAT   string `json:"created_at"`
	UpdatedAT   string `json:"updated_at,omitempty"`
	DeletedAT   string `json:"deleted_at,omitempty"`
}

type CustomerRequest struct {
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	PhoneNumber string `json:"phone_number"`
	Email       string `json:"email"`
	UserName    string `json:"user_name"`
	Password    string `json:"passward"`
	Gender      string `json:"gender"`
	Type        string `json:"type" binding:"required,oneof=superadmin user"`
	ImageURL    string `json:"image_url"`
}

type GetListCustomersResponse struct {
	Customers []*Customer `json:"customers"`
	Count     int64       `json:"count"`
}

type GetListParams struct {
	Limit  int64  `json:"limit" binding:"required" default:"10"`
	Page   int64  `json:"page" binding:"required" default:"1"`
	Search string `json:"search"`
}

type CustomerAuthResponse struct {
	Id          string `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	Email       string `json:"email,omitempty"`
	Username    string `json:"username,omitempty"`
	Gender      string `json:"gender,omitempty"`
	Type        string `json:"type,omitempty"`
	CreatedAt   string `json:"created_at,omitempty"`
	AccessToken string `json:"access_token,omitempty"`
}

type CreateCustomerRespons struct {
	ID    string `json:"id"`
	Token string `json:"token"`
}

func ParsCustomerRegisterToProtoStruct(customer *CustomerRequest) *pbu.CustomerRegisterRequest {
	return &pbu.CustomerRegisterRequest{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		UserName:    customer.UserName,
		Password:    customer.Password,
		Gender:      customer.Gender,
		Type:        customer.Type,
		ImageUrl:    customer.ImageURL,
	}
}

func ParsCustomerToProtoStruct(customer *CustomerRequest) *pbu.Customer {
	return &pbu.Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		UserName:    customer.UserName,
		Password:    customer.Password,
		Gender:      customer.Gender,
		Type:        customer.Type,
		ImageUrl:    customer.ImageURL,
	}
}

func ParsCustomerFromProtoStruct(customer *pbu.Customer) *Customer {
	return &Customer{
		ID:          customer.Id,
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		UserName:    customer.UserName,
		Password:    customer.Password,
		Gender:      customer.Gender,
		Type:        customer.Type,
		ImageURL:    customer.ImageUrl,
		CreatedAT:   customer.CreatedAt,
		UpdatedAT:   customer.UpdatedAt,
	}
}

func ParsListCustomersFromProtoStruct(customers []*pbu.Customer) (rCustomers []*Customer) {
	for _, cust := range customers {
		rCustomer := Customer{
			ID:          cust.Id,
			FirstName:   cust.FirstName,
			LastName:    cust.LastName,
			PhoneNumber: cust.PhoneNumber,
			Email:       cust.Email,
			UserName:    cust.UserName,
			Password:    cust.Password,
			Gender:      cust.Gender,
			Type:        cust.Type,
			ImageURL:    cust.ImageUrl,
			CreatedAT:   cust.CreatedAt,
			UpdatedAT:   cust.UpdatedAt,
		}
		rCustomers = append(rCustomers, &rCustomer)
	}
	return
}
