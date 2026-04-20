package handler

import (
	"go-project/internal/service"
	"go-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service: service}
}

func (h *userHandler) GetById(c *gin.Context) {
	user, err := h.service.GetById(c.Request.Context(), c.Param("id"))
	response.Response(c, user, err)
}

func (h *userHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll(c)
	response.Response(c, users, err)
}
