package handler

import (
	"go-project/internal/dto"
	"go-project/internal/service"
	"go-project/pkg/apperror"
	"go-project/pkg/response"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	GetById(c *gin.Context)
	GetAll(c *gin.Context)
	CreateUser(c *gin.Context)
	GetByEmail(c *gin.Context)
	Login(c *gin.Context)
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
	var user dto.CreateUserRequest
	if err := c.ShouldBindBodyWithJSON(&user); err != nil {
		response.Response(c, nil, apperror.ErrBadRequest.WithMessage(err.Error()))
		return
	}

	new_user := user.NewCreateUserRequest()
	result, err := h.service.CreateUser(c.Request.Context(), new_user)
	response.Response(c, result, err)
}

func (h *userHandler) Login(c *gin.Context) {
	var body map[string]string
	if err := c.ShouldBindJSON(&body); err != nil {
		response.Response(c, nil, apperror.ErrBadRequest.WithMessage(err.Error()))
		return
	}
	token, err := h.service.Login(c.Request.Context(), body["email"], body["password"])
	response.Response(c, token, err)
}
