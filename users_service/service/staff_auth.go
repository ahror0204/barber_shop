package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/barber_shop/users_service/config"
	pbu "github.com/barber_shop/users_service/genproto/users_service"
	emailPkg "github.com/barber_shop/users_service/pkg/email"

	l "github.com/barber_shop/users_service/pkg/logger"
	"github.com/barber_shop/users_service/pkg/utils"
	"github.com/barber_shop/users_service/storage"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	ErrWrongEmailOrPass = errors.New("wrong email or password")
	ErrEmailExists      = errors.New("email already exists")
	ErrUserNotVerified  = errors.New("user not verified")
	ErrIncorrectCode    = errors.New("incorrect verification code")
	ErrCodeExpired      = errors.New("verification code has been expired")
)

type StaffAuthService struct {
	pbu.UnimplementedStaffAuthServiceServer
	storage  storage.StorageI
	inMemory storage.InMemoryStorageI
	cfg      *config.Config
	logger   l.Logger
}

func NewStaffAuthService(s storage.StorageI, i storage.InMemoryStorageI, c *config.Config, l l.Logger) *StaffAuthService {
	return &StaffAuthService{
		storage:  s,
		inMemory: i,
		cfg:      c,
		logger:   l,
	}
}

func (s *StaffAuthService) StaffRegister(ctx context.Context, req *pbu.StaffRegisterRequest) (*pbu.Empty, error) {
	if err := emailPkg.ValidMailAddress(req.Email); err != nil {
		s.logger.Error("failed while validating email", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while validating email")
	}

	// verifing password
	if err := utils.VerifyPassword(req.Password); err != nil {
		s.logger.Error("failed while verifing password", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while verifing password")
	}

	// hashing password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		s.logger.Error("failed while hashing password", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while hashing password")
	}

	req.Password = hashedPassword

	staffData, err := json.Marshal(req)
	if err != nil {
		s.logger.Error("failed while marshaling staff", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while marshaling staff")
	}

	err = s.inMemory.Set("staff_"+req.Email, string(staffData), 10*time.Minute)
	if err != nil {
		s.logger.Error("failed to set to rd: %v", l.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to set to rd: %v", err)
	}

	go func() {
		err = s.sendVerificationCode(RegisterCodeKey, req.Email)
		if err != nil {
			fmt.Printf("failed to send verification code: %v", err)
		}
	}()

	return &pbu.Empty{}, nil
}

func (s *StaffAuthService) sendVerificationCode(key, email string) error {
	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		return err
	}
	err = s.inMemory.Set(key+email, code, time.Minute)
	if err != nil {
		return err
	}

	err = emailPkg.SendEmail(s.cfg, &emailPkg.SendEmailRequest{
		To:      []string{email},
		Subject: "Verification email",
		Body: map[string]string{
			"code": code,
		},
		Type: emailPkg.VerificationEmail,
	})
	if err != nil {
		return err
	}

	return nil
}

func (s *StaffAuthService) StaffVerify(ctx context.Context, req *pbu.VerifyStaffRegisterRequest) (*pbu.StaffAuthResponse, error) {
	staffData, err := s.inMemory.Get("staff_" + req.Email)
	if err != nil {
		s.logger.Error("failed while getting staff from redis", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting staff from redis")
	}
	var staff pbu.StaffRegisterRequest
	err = json.Unmarshal([]byte(staffData), &staff)
	if err != nil {
		s.logger.Error("failed while unmarshaling staff", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while unmarshaling staff")
	}

	code, err := s.inMemory.Get(RegisterCodeKey + req.Email)
	if err != nil {
		s.logger.Error("verification code has been expired", l.Error(err))
		return nil, status.Error(codes.Internal, "verification code has been expired")
	}

	if req.Code != code {
		s.logger.Error("incorrect code", l.Error(err))
		return nil, status.Error(codes.Internal, "incorrect code")
	}

	res, err := s.storage.Staff().CreateStaff(&pbu.Staff{
		SalonId:     staff.SalonId,
		FirstName:   staff.FirstName,
		LastName:    staff.LastName,
		PhoneNumber: staff.PhoneNumber,
		Email:       staff.Email,
		UserName:    staff.UserName,
		Password:    staff.Password,
		Type:        staff.Type,
		ImageUrl:    staff.ImageUrl,
	})
	if err != nil {
		s.logger.Error("failed while creating staff", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating staff")
	}

	//Creating token
	token, _, err := utils.CreateToken(s.cfg, &utils.TokenParams{
		UserID:   res.Id,
		Email:    res.Email,
		UserType: res.Type,
		Duration: time.Hour * 24,
	})
	if err != nil {
		s.logger.Error("failed while creating token", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating token")
	}

	return &pbu.StaffAuthResponse{
		Id:          res.Id,
		SalonId:     res.SalonId,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.UserName,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: token,
	}, nil
}

func (s *StaffAuthService) StaffLogIn(ctx context.Context, req *pbu.StaffLoginRequest) (*pbu.StaffAuthResponse, error) {

	res, err := s.storage.Staff().GetStaffByEmail(&pbu.Email{Email: req.Email})
	if err != nil {
		s.logger.Error("failed to get staff by email", l.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "staff not found: %v", ErrWrongEmailOrPass)
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	err = utils.CheckPassword(req.Password, res.Password)
	if err != nil {
		s.logger.Error("incorrect_password", l.Error(err))
		return nil, status.Error(codes.Internal, "incorrect_password")
	}

	//Creating token
	token, _, err := utils.CreateToken(s.cfg, &utils.TokenParams{
		UserID:   res.Id,
		Email:    res.Email,
		UserType: res.Type,
		Duration: time.Hour * 24,
	})
	if err != nil {
		s.logger.Error("failed while creating token", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating token")
	}

	return &pbu.StaffAuthResponse{
		Id:          res.Id,
		SalonId:     res.SalonId,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.UserName,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: token,
	}, nil
}

func (s *StaffAuthService) StaffForgotPassword(ctx context.Context, req *pbu.Email) (*pbu.Empty, error) {

	_, err := s.storage.Staff().GetStaffByEmail(&pbu.Email{Email: req.Email})
	if err != nil {
		s.logger.Error("failed to get staff by email", l.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "staff not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	go func() {
		err = s.sendVerificationCode(ForgotPasswordKey, req.Email)
		if err != nil {
			fmt.Printf("failed to send verification code: %v", err)
		}
	}()

	return &pbu.Empty{}, nil
}

func (s *StaffAuthService) VerifyStaffForgotPassword(ctx context.Context, req *pbu.VerifyStaffRegisterRequest) (*pbu.StaffAuthResponse, error) {

	code, err := s.inMemory.Get(ForgotPasswordKey + req.Email)
	if err != nil {
		s.logger.Error("verification code has been expired", l.Error(err))
		return nil, status.Error(codes.Internal, "verification code has been expired")
	}

	if req.Code != code {
		s.logger.Error("code comparison error", l.Error(err))
		return nil, status.Error(codes.Internal, "code comparison error")
	}

	res, err := s.storage.Staff().GetStaffByEmail(&pbu.Email{Email: req.Email})
	if err != nil {
		s.logger.Error("failed while getting staff by email", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting staff by email")
	}

	token, _, err := utils.CreateToken(s.cfg, &utils.TokenParams{
		UserID:   res.Id,
		Email:    res.Email,
		Duration: time.Minute * 30,
	})
	if err != nil {
		s.logger.Error("failed while creating token", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating token")
	}

	return &pbu.StaffAuthResponse{
		Id:          res.Id,
		SalonId:     res.SalonId,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.UserName,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: token,
	}, nil
}

func (s *StaffAuthService) UpdateStaffPassword(ctx context.Context, req *pbu.UpdatePasswordRequest) (*pbu.Empty, error) {
	// verifing password
	if err := utils.VerifyPassword(req.Password); err != nil {
		s.logger.Error("verify password error", l.Error(err))
		return nil, status.Error(codes.Internal, "verify password error")
	}

	// hashing password
	heshedPasword, err := utils.HashPassword(req.Password)
	if err != nil {
		s.logger.Error("hashing password error", l.Error(err))
		return nil, status.Error(codes.Internal, "hashing password error")
	}

	err = s.storage.Staff().UpdateStaffPassword(&pbu.UpdatePasswordRequest{
		ID:       req.ID,
		Password: heshedPasword,
	})
	if err != nil {
		s.logger.Error("error while updating staff password", l.Error(err))
		return nil, status.Error(codes.Internal, "error while updating staff password")
	}

	return &pbu.Empty{}, nil
}