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
	ID        string `json:"id"`
	UserID    string    `json:"customer_id"`
	Email     string    `json:"email"`
	Type      string    `json:"type"`
	IssuedAT  time.Time `json:"issued_at"`
	ExpiredAT time.Time `json:"expired_at"`
}

func NewPayload(params *TokenParams) (*Payload, error) {
	tokenID, err := uuid.NewRandom()
	if err != nil {
		return nil, err
	}

	return &Payload{
		ID:        tokenID.String(),
		UserID:    params.UserID,
		Email:     params.Email,
		Type:      params.Type,
		IssuedAT:  time.Now(),
		ExpiredAT: time.Now().Add(params.Duration),
	}, nil
}

// Valid checks if the token payload is valid or not
func (p *Payload) Valid() error {
	if time.Now().After(p.ExpiredAT) {
		return ErrExpiredToken
	}
	return nil
}
