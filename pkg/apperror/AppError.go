package apperror

type Kind string

type AppError struct {
	Code    int
	Message string
	Kind    Kind
}

const (
	NotFound     Kind = "not_found"
	Unauthorized Kind = "unauthorized"
	Conflict     Kind = "conflict"
)

func (e *AppError) Error() string { return e.Message }

var (
	ErrNotFound     = &AppError{Code: 404, Kind: NotFound, Message: "not found"}
	ErrUnauthorized = &AppError{Code: 401, Kind: Unauthorized, Message: "unauthorized"}
	ErrConflict     = &AppError{Code: 409, Kind: Conflict, Message: "conflict"}
)

func (e *AppError) Is(target error) bool {
	t, ok := target.(*AppError)
	if !ok {
		return false
	}
	if t.Message == "" {
		return e.Kind == t.Kind
	}
	return e.Kind == t.Kind && e.Message == t.Message
}
