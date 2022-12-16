package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/barber_shop/api-gateway/api/models"
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
	"github.com/gin-gonic/gin"
)

// @Router /staff/create [post]
// @Summary Create staff
// @Description This api for creating staff
// @Tags staff
// @Accept json
// @Produce json
// @Param staff body models.StaffRequest true "Staff"
// @Success 201 {object} models.Staff
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateStaff(c *gin.Context) {
	var req models.StaffRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	cstmr := models.ParsStaffToProtoStruct(&req)

	Staff, err := h.serviceManager.StaffService().CreateStaff(ctx, cstmr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ct := models.ParsStaffFromProtoStruct(Staff)

	c.JSON(http.StatusOK, ct)
}

// @Router /staff/update/{id} [put]
// @Summary Update a staff
// @Description This api for updating staff
// @Tags staff
// @Accept json
// @Produce json
// @Param id path string true "StaffID"
// @Param staff body models.UpdateStaffRequest true "Staff"
// @Success 200 {object} models.Staff
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateStaff(c *gin.Context) {
	var req models.UpdateStaffRequest
	id := c.Param("id")
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	staff := models.ParsUpdateStaffToProtoStruct(&req)
	staff.Id = id

	resp, err := h.serviceManager.StaffService().UpdateStaff(ctx, staff)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ParsStaffFromProtoStruct(resp))
}

// @Router /staff/get/{id} [get]
// @Summary Get staff by id
// @Description This api for getting staff by id
// @Tags staff
// @Accept json
// @Produce json
// @Param id path string true "StaffID"
// @Success 200 {object} models.Staff
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetStaffByID(c *gin.Context) {
	id := c.Param("id")

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	staff, err := h.serviceManager.StaffService().GetStaffByID(ctx, &pbu.ID{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ParsStaffFromProtoStruct(staff))
}

// @Router /staff/list [get]
// @Summary get list staff
// @Description This api for getting list of staff
// @Tags staff
// @Accept json
// @Produce json
// @Param filter query models.GetListParams false "Filter"
// @Success 200 {object} models.GetListStaffResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetListStaff(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.StaffService().GetListStaff(ctx, &pbu.GetListParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.GetListStaffResponse{
		Staff: models.ParsListStaffFromProtoStruct(res.Staff),
		Count: res.Count,
	})
}

// @Router /staff/delete/{id} [delete]
// @Summary Delete staff by id
// @Description This api for deleting staff by id
// @Tags staff
// @Accept json
// @Produce json
// @Param id path string true "StaffID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteStaff(c *gin.Context) {
	id := c.Param("id")

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err := h.serviceManager.StaffService().DeleteStaff(ctx, &pbu.ID{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "succesfuly deleted",
	})
}
