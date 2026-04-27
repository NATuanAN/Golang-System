package handler

import (
	"go-project/internal/model"
	"go-project/internal/service"
	"go-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	CreateUser(c *gin.Context)
	GetByEmail(c *gin.Context)
}

type userHandler struct {
	service service.UserService
}

func NewUserHandler(service service.UserService) UserHandler {
	return &userHandler{service}
}

func (h *userHandler) GetById(c *gin.Context) {
	user, err := h.service.GetById(c.Request.Context(), c.Param("id"))
	response.Response(c, user, err)
}

func (h *userHandler) GetAll(c *gin.Context) {
	users, err := h.service.GetAll(c.Request.Context())
	response.Response(c, users, err)
}
func (h *userHandler) GetByEmail(c *gin.Context) {
	user, err := h.service.GetByEmail(c.Request.Context(), c.Param("email"))
	response.Response(c, user, err)
}
func (h *userHandler) CreateUser(c *gin.Context) {
	var user model.User
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		response.Response(c, nil, err)
		return
	}
	result, err := h.service.CreateUser(c, &user)
	response.Response(c, result, err)
}
