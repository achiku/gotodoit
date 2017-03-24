package iapi

import "net/http"

type errCode string

// Error codes
const (
	InvalidParam        errCode = "invalid_param"
	Unauthorized                = "unauthorized"
	InternalServerError         = "internal_server_error"
	NotFound                    = "not_found"
)

// Error api errror
type Error struct {
	Code        errCode `json:"code,omitempty"`
	Description string  `json:"description,omitempty"`
}

// NewInternalServerError create internal server error
func NewInternalServerError() (int, Error) {
	return http.StatusInternalServerError, Error{
		Code:        InternalServerError,
		Description: "internal server error",
	}
}

// NewNotFoundError create internal server error
func NewNotFoundError() (int, Error) {
	return http.StatusInternalServerError, Error{
		Code:        NotFound,
		Description: "not found",
	}
}
