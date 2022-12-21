package v1

import (
	"context"

	"net/http"
	"time"

	"github.com/barber_shop/api-gateway/api/models"
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
	l "github.com/barber_shop/api-gateway/pkg/logger"
	"github.com/gin-gonic/gin"
)

// @Router /staff/register [post]
// @Summary staff register
// @Description This api for registeration staff
// @Tags staff_auth
// @Accept json
// @Produce json
// @Param data body models.StaffRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) StaffRegister(c *gin.Context) {
	var req models.StaffRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	//todo refactor get by email err
	res, _ := h.serviceManager.StaffService().GetStaffByEmail(ctx, &pbu.Email{Email: req.Email})
	if res != nil {
		c.JSON(http.StatusNotFound, errorResponse(ErrEmailExists))
		return
	}

	starReq := models.ParsStaffRegisterToProtoStruct(&req)

	_, err = h.serviceManager.StaffAuthService().StaffRegister(ctx, starReq)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("staff registration error", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Verification code has been sent!",
	})
}

// @Router /staff/verify [post]
// @Summary staff verify
// @Description This api for verification staff
// @Tags staff_auth
// @Accept json
// @Produse json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.StaffAuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) StaffVerify(c *gin.Context) {
	var req models.VerifyRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.StaffAuthService().StaffVerify(ctx, &pbu.VerifyStaffRegisterRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("staff verification error", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.StaffAuthResponse{
		Id:          res.Id,
		SalonId:     res.SalonId,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.Username,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: res.AccessToken,
	})
}

// @Router /staff/login [post]
// @Summary staff Login
// @Description This api for login staff
// @Tags staff_auth
// @Accept json
// @Produse json
// @Param data body models.LogInRequest true "Data"
// @Success 201 {object} models.StaffAuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) StaffLogIn(c *gin.Context) {
	var req models.LogInRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.StaffAuthService().StaffLogin(ctx, &pbu.StaffLoginRequest{Email: req.Email, Password: req.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("staff login error", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.StaffAuthResponse{
		Id:          res.Id,
		SalonId:     res.SalonId,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.Username,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: res.AccessToken,
	})
}

// @Router /staff/forgot-password [post]
// @Summary staff forgot password
// @Description This api for forgot password
// @Tags staff_auth
// @Accept json
// @Produce json
// @Param data body models.ForgotPasswordRequest true "Data"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) StaffForgotPassword(c *gin.Context) {
	var req models.ForgotPasswordRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err = h.serviceManager.StaffAuthService().StaffForgotPassword(ctx, &pbu.Email{Email: req.Email})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("forgot password func error", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Verification code has been sent!",
	})
}

// @Router /staff/verify-forgot-password [post]
// @Summary verify staff forgot password
// @Description Verify staff forgot password
// @Tags staff_auth
// @Accept json
// @Produce json
// @Param data body models.VerifyRequest true "Data"
// @Success 200 {object} models.StaffAuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) VerifyStaffForgotPassword(c *gin.Context) {
	var req models.VerifyRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		h.log.Error("failed while binding json", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.StaffAuthService().VerifyStaffForgotPassword(ctx, &pbu.VerifyStaffRegisterRequest{
		Email: req.Email,
		Code:  req.Code,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusCreated, models.StaffAuthResponse{
		Id:          res.Id,
		SalonId:     res.SalonId,
		FirstName:   res.FirstName,
		LastName:    res.LastName,
		Email:       res.Email,
		Username:    res.Username,
		Type:        res.Type,
		CreatedAt:   res.CreatedAt,
		AccessToken: res.AccessToken,
	})
}

// @Security ApiKeyAuth
// @Router /staff/update-password [post]
// @Summary staff update password
// @Description This api for updating staff password
// @Tags staff_auth
// @Accept json
// @Produce json
// @Param password body models.UpdatePasswordRequest true "Password"
// @Success 200 {object} models.StaffAuthResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateStaffPassword(c *gin.Context) {
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
		h.log.Error("failed while getting data from payload", l.Error(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err = h.serviceManager.StaffAuthService().UpdateStaffPassword(ctx, &pbu.UpdatePasswordRequest{
		ID:       payload.UserID,
		Password: req.Password,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		h.log.Error("failed while updating staff password", l.Error(err))
		return
	}

	c.JSON(http.StatusCreated, models.ResponseOK{
		Message: "Password has been updated!",
	})
}
