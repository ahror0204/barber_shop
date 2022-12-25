package service

import (
	"context"
	"time"

	pbu "github.com/barber_shop/users_service/genproto/users_service"
	"github.com/barber_shop/users_service/pkg/utils"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *StaffAuthService) VerifyToken(ctx context.Context, req *pbu.VerifyTokenRequest) (*pbu.AuthPayload, error) {
	accessTocken := req.AccessToken

	payload, err := utils.VerifyToken(*s.cfg, accessTocken)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
	}

	hasPermission, err := s.storage.Permission().CheckPermission(payload.UserType, req.Resource, req.Action)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "internal error: %v", err)
	}

	return &pbu.AuthPayload{
		Id:            payload.ID.String(),
		UserId:        payload.UserID,
		Email:         payload.Email,
		UserType:      payload.UserType,
		IssuedAt:      payload.IssuedAT.Format(time.RFC3339),
		ExpiredAt:     payload.ExpiredAT.Format(time.RFC3339),
		HasPermission: hasPermission,
	}, nil
}
