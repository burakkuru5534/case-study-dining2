// Package apperr defines domain-level errors that the HTTP layer maps to status codes.
package apperr

import (
	"errors"
	"net/http"
)

var (
	ErrNotFound   = errors.New("not found")
	ErrValidation = errors.New("validation failed")
	ErrConflict   = errors.New("conflict")
)

// HTTPStatus maps an error (possibly wrapped) to an HTTP status code.
func HTTPStatus(err error) int {
	switch {
	case errors.Is(err, ErrNotFound):
		return http.StatusNotFound
	case errors.Is(err, ErrValidation):
		return http.StatusBadRequest
	case errors.Is(err, ErrConflict):
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}
