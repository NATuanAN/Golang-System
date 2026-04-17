package domain

import "go-project/pkg/apperror"

var (
	ErrUserNotFound = &apperror.AppError{Code: 404, Kind: apperror.NotFound, Message: "user not found"}
	ErrUserExisted  = &apperror.AppError{Code: 409, Kind: apperror.Conflict, Message: "user already existed"}
)
