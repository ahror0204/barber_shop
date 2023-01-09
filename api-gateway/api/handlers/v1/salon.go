package v1

import (
	"context"
	"net/http"
	"time"

	"github.com/barber_shop/api-gateway/api/models"
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
	"github.com/gin-gonic/gin"
)

// @Security ApiKeyAuth
// @Router /salon/create [post]
// @Summary Create a salon
// @Description This api for creating salon
// @Tags salon
// @Accept json
// @Produce json
// @Param salon body models.SalonRequest true "Salon"
// @Success 201 {object} models.Salon
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateSalon(c *gin.Context) {
	var req models.SalonRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	salon := models.ParsSalonToProtoStruct(&req)

	resp, err := h.serviceManager.SalonService().CreateSalon(ctx, salon)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ParsSalonFromProtoStruct(resp))
}

// @Security ApiKeyAuth
// @Router /salon/update/{id} [put]
// @Summary Update salon
// @Description This api for updating salon
// @Tags salon
// @Accept json
// @Produce json
// @Param id path string true "SalonID"
// @Param salon body models.SalonRequest true "Salon"
// @Success 200 {object} models.Salon
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateSalon(c *gin.Context) {
	var req models.SalonRequest
	id := c.Param("id")
	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	salon := models.ParsSalonToProtoStruct(&req)
	salon.Id = id
	resp, err := h.serviceManager.SalonService().UpdateSalon(ctx, salon)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ParsSalonFromProtoStruct(resp))
}

// @Router /salon/get/{id} [get]
// @Summary Get salon by id
// @Description This api for getting salon by id
// @Tags salon
// @Accept json
// @Produce json
// @Param id path string true "SalonID"
// @Success 200 {object} models.Salon
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetSalonByID(c *gin.Context) {
	id := c.Param("id")

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	customer, err := h.serviceManager.SalonService().GetSalonByID(ctx, &pbu.ID{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ParsSalonFromProtoStruct(customer))
}

// @Router /salons/list [get]
// @Summary get list salons
// @Description This api for getting list of salons
// @Tags salon
// @Accept json
// @Produce json
// @Param filter query models.GetListParams false "Filter"
// @Success 200 {object} models.GetListSalonsResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetListSalons(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.SalonService().GetListSalons(ctx, &pbu.GetListParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.GetListSalonsResponse{
		Salons: models.ParsListSalonsFromProtoStruct(res.Salons),
		Count:  res.Count,
	})
}

// @Security ApiKeyAuth
// @Router /salon/delete/{id} [delete]
// @Summary Delete salon by id
// @Description This api for deleting salon by id
// @Tags salon
// @Accept json
// @Produce json
// @Param id path string true "SalonID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteSalon(c *gin.Context) {
	id := c.Param("id")

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err := h.serviceManager.SalonService().DeleteSalon(ctx, &pbu.ID{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "succesfuly deleted",
	})

}
