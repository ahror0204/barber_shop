package api

import (
	"github.com/barber_shop/users/storage"

	"github.com/gin-gonic/gin"
)

type handler struct{
	storage storage.StorageI
}

func NewService(s storage.StorageI) *gin.Engine {
	r := gin.Default()

	h := handler{
		storage: s,
	}

	r.POST("/create/user", h.CreateUser)
	r.PUT("/update/user", h.UpdateUser)
	r.GET("/get/user/:id", h.GetUser)
	r.GET("/allusers", h.GetAllUsers)
	r.DELETE("/delete/user/:id", h.DeleteUser)
	
	
	return r
}