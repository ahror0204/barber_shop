package utils

import (
	"errors"
	"time"

	"github.com/barber_shop/api-gateway/config"
	"github.com/golang-jwt/jwt"
)

const (
	UserTypeSuperadmin = "superadmin"
	UserTypeUser       = "user"
)

type TokenParams struct {
	UserID  string
	FirstName   string
	LastName    string
	Email       string
	UserName    string
	UserType    string
	UpdatedAT   string
	Duration    time.Duration
}

func CreateToken(cfg *config.Config, params *TokenParams) (string, *Payload, error) {
	payload, err := NewPayload(params)
	if err != nil {
		return "", payload, err
	}

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	token, err := jwtToken.SignedString([]byte(cfg.AuthSecretKey))
	return token, payload, err
}

func VerifyToken(cfg config.Config, token string) (*Payload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}
		return []byte(cfg.AuthSecretKey), nil
	}
	jwtToken, err := jwt.ParseWithClaims(token, &Payload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}
		return nil, ErrInvalidToken
	}

	paload, ok := jwtToken.Claims.(*Payload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return paload, nil
}
