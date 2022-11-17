package api

import (
	"net/http"
	"strconv"

	"github.com/barber_shop/users/structures"
	"github.com/gin-gonic/gin"
)

func (h *handler) CreateUser(c *gin.Context) {
	var user structures.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		errHandller(c, http.StatusBadRequest, err)
		return
	}

	id, err := h.storage.User().CreateUser(&user)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, id)

}

func (h *handler) UpdateUser(c *gin.Context) {
	var user structures.User
	err := c.ShouldBindJSON(&user)
	if err != nil {
		errHandller(c, http.StatusBadRequest, err)
		return
	}

	ruser, err := h.storage.User().UpdateUser(&user)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, ruser)

}

func (h *handler) GetUser(c *gin.Context) {
	id := c.Param("id")

	ruser, err := h.storage.User().GetUserByID(id)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(200, ruser)

}

func (h *handler) GetAllUsers(c *gin.Context) {
	params, err := validateGetAllParams(c)
	if err != nil {
		errHandller(c, http.StatusBadRequest, err)
		return
	}

	res, err := h.storage.User().GetAllUsers(params)
	if err != nil {
		errHandller(c, http.StatusInternalServerError, err)
		return
	}

	c.JSON(http.StatusOK, res)
}

func (h *handler) DeleteUser(c *gin.Context) {
	id := c.Param("id")

	err := h.storage.User().DeleteUser(id)
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

func validateGetAllParams(c *gin.Context) (*structures.GetUsersParams, error) {
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

	return &structures.GetUsersParams{
		Page: page,
		Limit: limit,
		Search: c.Query("search"),
	}, nil
}
