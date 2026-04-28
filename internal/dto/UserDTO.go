package dto

import "go-project/internal/model"

type CreateUserRequest struct {
	Name        string `json:"name"     binding:"required"`
	Email       string `json:"email"    binding:"required,email"`
	Password    string `json:"password" binding:"required,min=6"`
	PhoneNumber string `json:"phone_number"`
}

func (r *CreateUserRequest) NewCreateUserRequest() *model.User {
	return &model.User{
		Name:        r.Name,
		Email:       r.Email,
		Password:    r.Password,
		PhoneNumber: r.PhoneNumber,
		AccountType: model.AccountTypeBuyer,
		Active:      "active",
	}
}
