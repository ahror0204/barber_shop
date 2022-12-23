package v1

import (
	"context"
	"errors"
	"net/http"
	"time"

	"github.com/barber_shop/api-gateway/api/models"
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
	l "github.com/barber_shop/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

var (
	ErrWrongEmailOrPass = errors.New("wrong email or password")
	ErrEmailExists      = errors.New("email already exists")
	ErrUserNotVerified  = errors.New("user not verified")
	ErrIncorrectCode    = errors.New("incorrect verification code")
	ErrNotAllowed       = errors.New("method not allowed")
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
// @Tags customer_auth
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
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	//todo refactor get by email err
	res, err := h.serviceManager.CustomerService().GetCustomerByEmail(ctx, &pbu.Email{Email: req.Email})
	if res != nil {
		c.JSON(http.StatusNotFound, errorResponse(ErrEmailExists))
		h.log.Error("failed while getting customer by email", l.Error(err))
		return
	}

	customer := models.ParsCustomerRegisterToProtoStruct(&req)

	_, err = h.serviceManager.CustomerAuthService().CustomerRegister(ctx, customer)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("customer registration error", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Verification code has been sent!",
	})
}

// @Router /customer/verify [post]
// @Summary Verify customer
// @Description This api for verification customer
// @Tags customer_auth
// @Accept json
// @Produse json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.CustomerAuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CustomerVerify(c *gin.Context) {
	var req models.VerifyRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.CustomerAuthService().CustomerVerify(ctx, &pbu.VerifyCustomerRegisterRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("customer verification error", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.CustomerAuthResponse{
		Id:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.Username,
		Gender:      res.Gender,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: res.AccessToken,
	})
}

// @Router /customer/login [post]
// @Summary Login customer
// @Description This api for login customer
// @Tags customer_auth
// @Accept json
// @Produse json
// @Param data body models.LogInRequest true "Data"
// @Success 201 {object} models.CustomerAuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CustomerLogIn(c *gin.Context) {
	var req models.LogInRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.CustomerAuthService().CustomerLogin(ctx, &pbu.CustomerLoginRequest{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("staff login error", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.CustomerAuthResponse{
		Id:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.Username,
		Gender:      res.Gender,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: res.AccessToken,
	})
}

// @Router /customer/forgot-password [post]
// @Summary forgot password
// @Description This api for forgot password
// @Tags customer_auth
// @Accept json
// @Produce json
// @Param data body models.ForgotPasswordRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CustomerForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err = h.serviceManager.CustomerAuthService().CustomerForgotPassword(ctx, &pbu.Email{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("customer forgot password error", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Verification code has been sent!",
	})
}

// @Router /customer/verify-forgot-password [post]
// @Summary Verify forgot password
// @Description Verify forgot password
// @Tags customer_auth
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) VerifyForgotPassword(c *gin.Context) {
	var req models.VerifyRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.CustomerAuthService().VerifyCustomerForgotPassword(ctx, &pbu.VerifyCustomerRegisterRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("customer login error", l.Error(err))
		return
	}

	c.JSON(http.StatusOK, models.CustomerAuthResponse{
		Id:          res.Id,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.Username,
		Gender:      res.Gender,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: res.AccessToken,
	})
}

// @Security ApiKeyAuth
// @Router /customer/update-password [post]
// @Summary update password
// @Description This api for updating customer password
// @Tags customer_auth
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
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("failed while getting customer from payload", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err = h.serviceManager.CustomerAuthService().UpdateCustomerPassword(ctx, &pbu.UpdatePasswordRequest{
		ID:       payload.UserID,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("failed while updating customer password", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Password has been updated!",
	})
}
