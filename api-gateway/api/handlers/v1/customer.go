package v1

import (
	"context"
	"net/http"

	"github.com/barber_shop/api-gateway/api/models"
	pbu "github.com/barber_shop/api-gateway/genproto/users_service"
	"github.com/gin-gonic/gin"
	"time"
)

// @Security ApiKeyAuth
// @Router /customer/create [post]
// @Summary Create a customer
// @Description This api for creating customer
// @Tags customer
// @Accept json
// @Produce json
// @Param customer body models.CustomerRequest true "Customer"
// @Success 201 {object} models.CreateCustomerRespons
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateCustomer(c *gin.Context) {
	var req models.CustomerRequest

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	cstmr := models.ParsCustomerToProtoStruct(&req)

	customer, err := h.serviceManager.CustomerService().CreateCustomer(ctx, cstmr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ct := models.ParsCustomerFromProtoStruct(customer)

	c.JSON(http.StatusOK, ct)
}

// @Security ApiKeyAuth
// @Router /customer/update/{id} [put]
// @Summary Update a customer
// @Description This api for updating customer
// @Tags customer
// @Accept json
// @Produce json
// @Param id path string true "CustomerID"
// @Param customer body models.CustomerRequest true "Customer"
// @Success 200 {object} models.Customer
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) UpdateCustomer(c *gin.Context) {
	var customer models.CustomerRequest
	id := c.Param("id")
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	cstmr := models.ParsCustomerToProtoStruct(&customer)
	cstmr.Id = id
	rCustomer, err := h.serviceManager.CustomerService().UpdateCustomer(ctx, cstmr)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ParsCustomerFromProtoStruct(rCustomer))
}

// @Router /customer/get/{id} [get]
// @Summary Get customer by id
// @Description This api for getting customer by id
// @Tags customer
// @Accept json
// @Produce json
// @Param id path string true "CustomerID"
// @Success 200 {object} models.Customer
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCustomerByID(c *gin.Context) {
	id := c.Param("id")

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	customer, err := h.serviceManager.CustomerService().GetCustomerByID(ctx, &pbu.ID{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ParsCustomerFromProtoStruct(customer))
}

// @Security ApiKeyAuth
// @Router /customer/me [get]
// @Summary Get customer by token
// @Description This api for getting customer by token
// @Tags customer
// @Accept json
// @Produce json
// @Success 200 {object} models.Customer
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetCustomerProfile(c *gin.Context) {
	payload, err := h.GetAuthPayload(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	resp, err := h.serviceManager.CustomerService().GetCustomerByID(ctx, &pbu.ID{Id: payload.UserID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ParsCustomerFromProtoStruct(resp))
}

// @Router /customers/list [get]
// @Summary get list customers
// @Description This api for getting list of customers
// @Tags customer
// @Accept json
// @Produce json
// @Param filter query models.GetListParams false "Filter"
// @Success 200 {object} models.GetListCustomersResponse
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) GetListCustomers(c *gin.Context) {
	req, err := validateGetAllParams(c)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	res, err := h.serviceManager.CustomerService().GetListCustomers(ctx, &pbu.GetListParams{
		Page:   req.Page,
		Limit:  req.Limit,
		Search: req.Search,
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.GetListCustomersResponse{
		Customers: models.ParsListCustomersFromProtoStruct(res.Customers),
		Count:     res.Count,
	})
}

// @Security ApiKeyAuth
// @Router /customer/delete/{id} [delete]
// @Summary Delete customer by id
// @Description This api for deleting customer by id
// @Tags customer
// @Accept json
// @Produce json
// @Param id path string true "CustomerID"
// @Success 200 {object} models.ResponseOK
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) DeleteCustomer(c *gin.Context) {
	id := c.Param("id")

	ctx, cencel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cencel()

	_, err := h.serviceManager.CustomerService().DeleteCustomer(ctx, &pbu.ID{Id: id})
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	c.JSON(http.StatusOK, models.ResponseOK{
		Message: "succesfuly deleted",
	})

}
