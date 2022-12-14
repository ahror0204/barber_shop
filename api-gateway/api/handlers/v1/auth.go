package v1

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/barber_shop/api-gateway/api/models"
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
	emailPkg "github.com/barber_shop/api-gateway/pkg/email"
	l "github.com/barber_shop/api-gateway/pkg/logger"
	"github.com/barber_shop/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

var (
	ErrWrongEmailOrPass = errors.New("wrong email or password")
	ErrEmailExists      = errors.New("email already exists")
	ErrUserNotVerified  = errors.New("user not verified")
	ErrIncorrectCode    = errors.New("incorrect verification code")
	ErrCodeExpired      = errors.New("verification code has been expired")
)

const (
	RegisterCodeKey   = "register_code_"
	ForgotPasswordKey = "forgot_password_code_"
	expireTimeSecond  = 600
)

// @Router /customer/register [post]
// @Summary register a customer
// @Description This api for registering a customer
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.CustomerRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) RegisterCustomer(c *gin.Context) {
	var req models.CustomerRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	if err := emailPkg.ValidMailAddress(req.Email); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(ErrWrongEmailOrPass))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	//todo reactor get by email err
	res, err := h.serviceManager.CustomerService().GetCustomerByEmail(ctx, &pbu.Email{Email: req.Email})
	if res != nil {
		c.JSON(http.StatusNotFound, errorResponse(ErrEmailExists))
		return
	}

	// verifing password
	if err := utils.VerifyPassword(req.Password); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("password verify error", l.Error(err))
		return
	}

	// hashing password
	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	req.Password = hashedPassword

	customer := models.ParsCustomerToProtoStruct(&req)

	customerData, err := json.Marshal(customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = h.redisStorage.SetWithTTL(customer.Email, string(customerData), expireTimeSecond)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	go func() {
		err = h.sendVerificationCode(RegisterCodeKey, req.Email)
		if err != nil {
			fmt.Printf("failed to send verification code: %v", err)
		}
	}()

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Verification code has been sent!",
	})

}

func (h *handlerV1) sendVerificationCode(key, email string) error {
	code, err := utils.GenerateRandomCode(6)
	if err != nil {
		return err
	}

	err = h.redisStorage.SetWithTTL(key+email, code, 60)
	if err != nil {
		return err
	}

	err = emailPkg.SendEmail(&h.cfg, &emailPkg.SendEmailRequest{
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

// @Router /customer/verify [post]
// @Summary Verify customer
// @Description This api for verification customer
// @Tags auth
// @Accept json
// @Produse json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.CreateCustomerRespons
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) Verify(c *gin.Context) {
	var (
		req      models.VerifyRequest
		customer pbu.Customer
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	body, err := h.redisStorage.Get(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = json.Unmarshal(body.([]byte), &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	codeRedis, err := h.redisStorage.Get(RegisterCodeKey + req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(ErrCodeExpired))
		return
	}

	code := string((codeRedis.([]byte))[:])

	if req.Code != code {
		c.JSON(http.StatusForbidden, errorResponse(ErrIncorrectCode))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	id, err := h.serviceManager.CustomerService().CreateCustomer(ctx, &customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	//Creating token
	token, _, err := utils.CreateToken(&h.cfg, &utils.TokenParams{
		CustomerID: id.Id,
		Duration:   time.Hour * 24,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.CreateCustomerRespons{ID: id.Id, Token: token})
}

// @Router /customer/login [post]
// @Summary Login customer
// @Description This api for login customer
// @Tags auth
// @Accept json
// @Produse json
// @Param data body models.LogInCustomerRequest true "Data"
// @Success 201 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CustomerLogIn(c *gin.Context) {
	var (
		req models.LogInCustomerRequest
	)

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	result, err := h.serviceManager.CustomerService().GetCustomerByEmail(ctx, &pbu.Email{Email: req.Email})
	if err != nil {
		if result == nil {
			c.JSON(http.StatusForbidden, errorResponse(ErrWrongEmailOrPass))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = utils.CheckPassword(req.Password, result.Password)
	if err != nil {
		c.JSON(http.StatusForbidden, errorResponse(ErrWrongEmailOrPass))
		return
	}

	// creating token
	token, _, err := utils.CreateToken(&h.cfg, &utils.TokenParams{
		CustomerID:  result.Id,
		FirstName:   result.FirstName,
		LastName:    result.LastName,
		PhoneNumber: result.PhoneNumber,
		Email:       result.Email,
		UserName:    result.UserName,
		Password:    result.Password,
		Gender:      result.Gender,
		ImageURL:    result.ImageUrl,
		CreatedAT:   result.CreatedAt,
		UpdatedAT:   result.UpdatedAt,
		Duration:    time.Hour * 24,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := models.ParsAuthResponseToPbCustomer(result)

	resp.Token = token

	c.JSON(http.StatusCreated, resp)
}

// @Router /customer/forgot-password [post]
// @Summary forgot password
// @Description This api for forgot password
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.ForgotPasswordRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) ForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	resp, err := h.serviceManager.CustomerService().GetCustomerByEmail(ctx, &pbu.Email{Email: req.Email})
	if err != nil {
		if resp == nil {
			c.JSON(http.StatusNotFound, errorResponse(err))
			return
		}

		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	go func() {
		err = h.sendVerificationCode(ForgotPasswordKey, req.Email)
		if err != nil {
			fmt.Printf("failed to send verification code: %v", err)
		}
	}()

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Verification code has been sent!",
	})
}

// @Router /customer/verify-forgot-password [post]
// @Summary Verify forgot password
// @Description Verify forgot password
// @Tags auth
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) VerifyForgotPassword(c *gin.Context) {
	var req models.VerifyRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	data, err := h.redisStorage.Get(ForgotPasswordKey + req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(ErrCodeExpired))
		return
	}

	// todo refactor (gave panic when code expired)
	code := string((data.([]byte))[:])
	if req.Code != code {
		c.JSON(http.StatusForbidden, errorResponse(ErrIncorrectCode))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.CustomerService().GetCustomerByEmail(ctx, &pbu.Email{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	token, _, err := utils.CreateToken(&h.cfg, &utils.TokenParams{
		CustomerID:  res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		PhoneNumber: res.PhoneNumber,
		Email:       res.Email,
		UserName:    res.UserName,
		Password:    res.Password,
		Gender:      res.Gender,
		ImageURL:    res.ImageUrl,
		CreatedAT:   res.CreatedAt,
		UpdatedAT:   res.UpdatedAt,
		Duration:    time.Minute * 30,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	response := models.ParsAuthResponseToPbCustomer(res)
	response.Token = token

	c.JSON(http.StatusCreated, response)
}

// @Security ApiKeyAuth
// @Router /customer/update-password [post]
// @Summary update password
// @Description This api for updating customer password
// @Tags auth
// @Accept json
// @Produce json
// @Param password body models.UpdatePasswordRequest true "Password"
// @Success 200 {object} models.AuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateCustomerPassword(c *gin.Context) {
	var req models.UpdatePasswordRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	// verifing password
	if err := utils.VerifyPassword(req.Password); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("password verify error", l.Error(err))
		return
	}

	// hashing password
	heshedPasword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err = h.serviceManager.CustomerService().UpdateCustomerPassword(ctx, &pbu.UpdateCustomerPasswordRequest{
		ID:       payload.CustomerID,
		Password: heshedPasword,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Password has been updated!",
	})
}
