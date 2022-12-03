package utils

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

// Different types of error returned by the VerifyToken function
var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type Payload struct {
	ID          uuid.UUID `json:"id"`
	CustomerID  string    `json:"customer_id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	PhoneNumber string    `json:"phone_number"`
	Email       string    `json:"email"`
	UserName    string    `json:"user_name"`
	Password    string    `json:"password"`
	Gender      string    `json:"gender"`
	ImageURL    string    `json:"image_url"`
	CreatedAT   string    `json:"created_at"`
	UpdatedAT   string    `json:"updated_at"`
	IssuedAT    time.Time `json:"issued_at"`
	ExpiredAT   time.Time `json:"expired_at"`
}

func NewPayload(params *TokenParams) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:         tokenID,
		CustomerID:  params.CustomerID,
		FirstName:   params.FirstName,
		LastName:    params.LastName,
		PhoneNumber: params.PhoneNumber,
		Email:       params.Email,
		UserName:    params.UserName,
		Password:    params.Password,
		Gender:      params.Gender,
		ImageURL:    params.ImageURL,
		CreatedAT:   params.CreatedAT,
		UpdatedAT:   params.UpdatedAT,
		IssuedAT:   time.Now(),
		ExpiredAT:  time.Now().Add(params.Duration),
	}, nil
}

// Valid checks if the token payload is valid or not
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAT) {
		return ErrExpiredToken
	}
	return nil
}
