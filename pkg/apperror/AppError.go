package apperror

type Kind string

type AppError struct {
	code    int
	message string
	kind    Kind
}

const (
	NotFound     Kind = "not_found"
	Unauthorized Kind = "unauthorized"
	Conflict     Kind = "conflict"
	BadRequest   Kind = "bad_request"
)

func (e *AppError) Error() string { return e.message }
func (e *AppError) Code() int     { return e.code }
func (e *AppError) Kind() Kind    { return e.kind }

var (
	ErrNotFound     = &AppError{code: 404, kind: NotFound, message: "not found"}
	ErrUnauthorized = &AppError{code: 401, kind: Unauthorized, message: "unauthorized"}
	ErrConflict     = &AppError{code: 409, kind: Conflict, message: "conflict"}
	ErrBadRequest   = &AppError{code: 400, kind: BadRequest, message: "bad_request"}
)

func (e *AppError) WithMessage(message string) *AppError {
	return &AppError{
		code:    e.code,
		kind:    e.kind,
		message: message,
	}
}
func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	return e.kind == t.kind
}
