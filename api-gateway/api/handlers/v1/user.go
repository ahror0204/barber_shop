package v1

import (
	"context"
	"net/http"

	"time"
	pb "github.com/barber_shop/api-gateway/genproto"
	"github.com/barber_shop/api-gateway/api/models"
	"github.com/gin-gonic/gin"
)

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
	var customer models.CustomerRequest
	
	err := c.ShouldBindJSON(&customer)
	if err != nil {
		errHandller(c, http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	cstmr := models.ParsCustomerToProtoStruct(&customer)

	ID, err := h.serviceManager.UserService().CreateCustomer(ctx, cstmr)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, models.CreateCustomerRespons{
		ID: ID.Id,
	})
}

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
		errHandller(c, http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	cstmr := models.ParsCustomerToProtoStruct(&customer)
	cstmr.Id = id
	rCustomer, err := h.serviceManager.UserService().UpdateCustomer(ctx, cstmr)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	res := models.ParsCustomerFromProtoStruct(rCustomer)

	c.JSON(http.StatusOK, res)
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

	customer, err := h.serviceManager.UserService().GetCustomerByID(ctx, &pb.ID{Id: id})
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	res := models.ParsCustomerFromProtoStruct(customer)

	c.JSON(http.StatusOK, res)
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

	res, err := h.serviceManager.UserService().GetListCustomers(ctx, &pb.GetCustomerParams{
		Page: req.Page,
		Limit: req.Limit,
		Search: req.Search,
	})
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	customers := models.ParsListCustomersFromProtoStruct(res.Customers)

	c.JSON(http.StatusOK, models.GetListCustomersResponse{
		Customers: customers,
		Count: res.Count,
	})
}

// func (h *handlerV1) DeleteUser(c *gin.Context) {
// 	id := c.Param("id")

// 	_, err := h.serviceManager.UserService().DeleteUser(id)
// 	if err != nil {
// 		errHandller(c, http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(200, "successful")

// }

func errHandller(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

// func validateGetAllParams(c *gin.Context) (*models.GetUsersParams, error) {
// 	var (
// 		limit int = 10
// 		page  int = 1
// 		err   error
// 	)

// 	if c.Query("limit") != "" {
// 		limit, err = strconv.Atoi(c.Query("limit"))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	if c.Query("page") != "" {
// 		page, err = strconv.Atoi(c.Query("page"))
// 		if err != nil {
// 			return nil, err
// 		}
// 	}

// 	return &models.GetUsersParams{
// 		Page:   page,
// 		Limit:  limit,
// 		Search: c.Query("search"),
// 	}, nil
// }
