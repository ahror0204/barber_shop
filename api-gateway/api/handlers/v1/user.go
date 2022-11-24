package v1

import (
	"context"
	"net/http"

	"time"

	"github.com/barber_shop/api-gateway/api/models"
	"github.com/gin-gonic/gin"
)


// @Router /user/create [post]
// @summary Create a user
// @Description This api for creating user
// @Tags user
// @Accept json
// @Produce json
// @Param user body models.CreateUserRequest true "User"
// @Success 201 {object} models.User
// @Failure 500 {object} models.ErrorResponse
func (h *handlerV1) CreateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		errHandller(c, http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()
	
	puser := models.ParsUserStruct(user)
	
	ID, err := h.serviceManager.UserService().CreateUser(ctx, &puser)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, models.CreateUserRespons{
		ID: ID.Id,
	})
}

func (h *handlerV1) UpdateUser(c *gin.Context) {
	var user models.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		errHandller(c, http.StatusBadRequest, err)
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*time.Duration(h.cfg.CtxTimeout))
	defer cancel()

	puser := models.ParsUserStruct(user)

	ruser, err := h.serviceManager.UserService().UpdateUser(ctx, &puser)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, ruser)
}

func (h *handlerV1) GetUser(c *gin.Context) {
	id := c.Param("id")
	ruser, err := h.serviceManager.UserService().GetUserByID(id)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, ruser)

}

func (h *handlerV1) GetAllUsers(c *gin.Context) {
	params, err := validateGetAllParams(c)
	if err != nil {
		errHandller(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.serviceManager.UserService().GetAllUsers(params)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handlerV1) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	_, err := h.serviceManager.UserService().DeleteUser(id)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, "successful")

}

func errHandller(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"error": err.Error(),
	})
}

func validateGetAllParams(c *gin.Context) (*models.GetUsersParams, error) {
	var (
		limit int = 10
		page  int = 1
		err   error
	)

	if c.Query("limit") != "" {
		limit, err = strconv.Atoi(c.Query("limit"))
		if err != nil {
			return nil, err
		}
	}

	if c.Query("page") != "" {
		page, err = strconv.Atoi(c.Query("page"))
		if err != nil {
			return nil, err
		}
	}

	return &models.GetUsersParams{
		Page:   page,
		Limit:  limit,
		Search: c.Query("search"),
	}, nil
}
 