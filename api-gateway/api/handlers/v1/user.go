package v1

import (
	"context"
	"fmt"
	"net/http"

	"time"

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

// @Router /customer/update [put]
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
	fmt.Println(id, cstmr, "----------------------------------------")
	rCustomer, err := h.serviceManager.UserService().UpdateCustomer(ctx, cstmr)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	res := models.ParsCustomerFromProtoStruct(rCustomer)

	c.JSON(http.StatusOK, res)
}

// func (h *handlerV1) GetUser(c *gin.Context) {
// 	id := c.Param("id")
// 	ruser, err := h.serviceManager.UserService().GetUserByID(id)
// 	if err != nil {
// 		errHandller(c, http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(200, ruser)

// }

// func (h *handlerV1) GetAllUsers(c *gin.Context) {
// 	params, err := validateGetAllParams(c)
// 	if err != nil {
// 		errHandller(c, http.StatusBadRequest, err)
// 		return
// 	}

// 	res, err := h.serviceManager.UserService().GetAllUsers(params)
// 	if err != nil {
// 		errHandller(c, http.StatusInternalServerError, err)
// 		return
// 	}

// 	c.JSON(http.StatusOK, res)
// }

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
