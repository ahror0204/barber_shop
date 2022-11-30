package v1

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/barber_shop/api-gateway/api/models"
	pb "github.com/barber_shop/api-gateway/genproto"
	emailPkg "github.com/barber_shop/api-gateway/pkg/email"
	"github.com/barber_shop/api-gateway/pkg/utils"
	"github.com/gin-gonic/gin"
)

var (
	ErrEmailExists = errors.New("email already exists")
)

const (
	RegisterCodeKey   = "register_code_"
	ForgotPasswordKey = "forgot_password_code_"
)

// @Router /customer/register [post]
// @Summary register a customer
// @Description This api for registering a customer
// @Tags auth
// @Accept json
// @Produce json
// @Param data body modules.RegisterCustomer true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) RegisterCustomer(c *gin.Context) {
	var req models.CustomerRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err = h.serviceManager.UserService().GetCustomerByEmail(ctx, &pb.Email{Email: req.Email})

	if !errors.Is(err, sql.ErrNoRows) {
		c.JSON(http.StatusBadRequest, errorResponse(ErrEmailExists))
		return
	}

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

	err = h.redisStorage.SetWithTTL(customer.Email, string(customerData), 10*time.Minute)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	go func ()  {
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

	err = h.redisStorage.SetWithTTL(key+email, code, time.Minute)
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
