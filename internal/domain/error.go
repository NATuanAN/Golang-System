package domain

import "go-project/pkg/apperror"

type err interface {
	Code() int
	Kind() apperror.Kind
	Error() string
}
