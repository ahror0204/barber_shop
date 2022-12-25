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

type CustomerAuthService struct {
	pbu.UnimplementedCustomerAuthServiceServer
	storage  storage.StorageI
	inMemory storage.InMemoryStorageI
	cfg      *config.Config
	logger   l.Logger
}

func NewCustomerAuthService(s storage.StorageI, inMemory storage.InMemoryStorageI, cfg *config.Config, logger l.Logger) *CustomerAuthService {
	return &CustomerAuthService{
		storage:  s,
		inMemory: inMemory,
		cfg:      cfg,
		logger:   logger,
	}
}

const (
	RegisterCodeKey   = "register_code_"
	ForgotPasswordKey = "forgot_password_code_"
)

func (c *CustomerAuthService) CustomerRegister(ctx context.Context, req *pbu.CustomerRegisterRequest) (*pbu.Empty, error) {
	if err := emailPkg.ValidMailAddress(req.Email); err != nil {
		c.logger.Error("failed while validating email", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while validating email")
	}

	// verifing password
	if err := utils.VerifyPassword(req.Password); err != nil {
		c.logger.Error("failed while verifing password", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while verifing password")
	}

	// hashing password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.logger.Error("failed while hashing password", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while hashing password")
	}

	req.Password = hashedPassword

	customerData, err := json.Marshal(req)
	if err != nil {
		c.logger.Error("failed while marshaling customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while marshaling customer")
	}

	err = c.inMemory.Set("customer_"+req.Email, string(customerData), 10*time.Minute)
	if err != nil {
		c.logger.Error("failed to set to rd: %v", l.Error(err))
		return nil, status.Errorf(codes.Internal, "failed to set to rd: %v", err)
	}

	go func() {
		err = c.sendVerificationCode(RegisterCodeKey, req.Email)
		if err != nil {
			fmt.Printf("failed to send verification code: %v", err)
		}
	}()

	return &pbu.Empty{}, nil
}

func (c *CustomerAuthService) sendVerificationCode(key, email string) error {
	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		return err
	}
	err = c.inMemory.Set(key+email, code, time.Minute)
	if err != nil {
		return err
	}

	err = emailPkg.SendEmail(c.cfg, &emailPkg.SendEmailRequest{
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

func (c *CustomerAuthService) CustomerVerify(ctx context.Context, req *pbu.VerifyCustomerRegisterRequest) (*pbu.CustomerAuthResponse, error) {
	customerData, err := c.inMemory.Get("customer_" + req.Email)
	if err != nil {
		c.logger.Error("failed while getting customer from redis", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting customer from redis")
	}
	var customer pbu.CustomerRegisterRequest
	err = json.Unmarshal([]byte(customerData), &customer)
	if err != nil {
		c.logger.Error("failed while unmarshaling customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while unmarshaling customer")
	}

	code, err := c.inMemory.Get(RegisterCodeKey + req.Email)
	if err != nil {
		c.logger.Error("verification code has been expired", l.Error(err))
		return nil, status.Error(codes.Internal, "verification code has been expired")
	}

	if req.Code != code {
		c.logger.Error("incorrect code", l.Error(err))
		return nil, status.Error(codes.Internal, "incorrect code")
	}

	res, err := c.storage.Customer().CreateCustomer(&pbu.Customer{
		FirstName:   customer.FirstName,
		LastName:    customer.LastName,
		PhoneNumber: customer.PhoneNumber,
		Email:       customer.Email,
		UserName:    customer.UserName,
		Password:    customer.Password,
		Gender:      customer.Gender,
		Type:        customer.Type,
		ImageUrl:    customer.ImageUrl,
	})
	if err != nil {
		c.logger.Error("failed while creating customer", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating customer")
	}

	//Creating token
	token, _, err := utils.CreateToken(c.cfg, &utils.TokenParams{
		UserID:   res.Id,
		Email:    res.Email,
		UserType: res.Type,
		Duration: time.Hour * 24,
	})
	if err != nil {
		c.logger.Error("failed while creating token", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating token")
	}

	return &pbu.CustomerAuthResponse{
		Id:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.UserName,
		Gender:      res.Gender,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: token,
	}, nil
}

func (c *CustomerAuthService) CustomerLogin(ctx context.Context, req *pbu.CustomerLoginRequest) (*pbu.CustomerAuthResponse, error) {

	res, err := c.storage.Customer().GetCustomerByEmail(&pbu.Email{Email: req.Email})
	if err != nil {
		c.logger.Error("failed to get customer by email", l.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "customer not found: %v", ErrWrongEmailOrPass)
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	err = utils.CheckPassword(req.Password, res.Password)
	if err != nil {
		c.logger.Error("incorrect_password", l.Error(err))
		return nil, status.Error(codes.Internal, "incorrect_password")
	}

	//Creating token
	token, _, err := utils.CreateToken(c.cfg, &utils.TokenParams{
		UserID:   res.Id,
		Email:    res.Email,
		UserType: res.Type,
		Duration: time.Hour * 24,
	})
	if err != nil {
		c.logger.Error("failed while creating token", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating token")
	}

	return &pbu.CustomerAuthResponse{
		Id:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.UserName,
		Gender:      res.Gender,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: token,
	}, nil
}

func (c *CustomerAuthService) CustomerForgotPassword(ctx context.Context, req *pbu.Email) (*pbu.Empty, error) {

	_, err := c.storage.Customer().GetCustomerByEmail(&pbu.Email{Email: req.Email})
	if err != nil {
		c.logger.Error("failed to get customer by email", l.Error(err))
		if errors.Is(err, sql.ErrNoRows) {
			return nil, status.Errorf(codes.NotFound, "customer not found: %v", err)
		}
		return nil, status.Errorf(codes.Internal, "internal server error: %v", err)
	}

	go func() {
		err = c.sendVerificationCode(ForgotPasswordKey, req.Email)
		if err != nil {
			fmt.Printf("failed to send verification code: %v", err)
		}
	}()

	return &pbu.Empty{}, nil
}

func (c *CustomerAuthService) VerifyCustomerForgotPassword(ctx context.Context, req *pbu.VerifyCustomerRegisterRequest) (*pbu.CustomerAuthResponse, error) {

	code, err := c.inMemory.Get(ForgotPasswordKey + req.Email)
	if err != nil {
		c.logger.Error("verification code has been expired", l.Error(err))
		return nil, status.Error(codes.Internal, "verification code has been expired")
	}

	if req.Code != code {
		c.logger.Error("code comparison error", l.Error(err))
		return nil, status.Error(codes.Internal, "code comparison error")
	}

	res, err := c.storage.Customer().GetCustomerByEmail(&pbu.Email{Email: req.Email})
	if err != nil {
		c.logger.Error("failed while getting customer by email", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while getting customer by email")
	}

	token, _, err := utils.CreateToken(c.cfg, &utils.TokenParams{
		UserID:   res.Id,
		Email:    res.Email,
		Duration: time.Minute * 30,
	})
	if err != nil {
		c.logger.Error("failed while creating token", l.Error(err))
		return nil, status.Error(codes.Internal, "failed while creating token")
	}

	return &pbu.CustomerAuthResponse{
		Id:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.UserName,
		Gender:      res.Gender,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: token,
	}, nil
}

func (c *CustomerAuthService) UpdateCustomerPassword(ctx context.Context, req *pbu.UpdatePasswordRequest) (*pbu.Empty, error) {

	// verifing password
	if err := utils.VerifyPassword(req.Password); err != nil {
		c.logger.Error("verify password error", l.Error(err))
		return nil, status.Error(codes.Internal, "verify password error")
	}

	// hashing password
	hashedPasword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.logger.Error("hashing password error", l.Error(err))
		return nil, status.Error(codes.Internal, "hashing password error")
	}

	err = c.storage.Customer().UpdateCustomerPassword(&pbu.UpdatePasswordRequest{
		ID:       req.ID,
		Password: hashedPasword,
	})
	if err != nil {
		c.logger.Error("error while updating customer password", l.Error(err))
		return nil, status.Error(codes.Internal, "error while updating customer password")
	}

	return &pbu.Empty{}, nil
}
