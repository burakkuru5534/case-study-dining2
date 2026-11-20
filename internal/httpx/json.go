// Package httpx holds small HTTP helpers shared by handlers.
package httpx

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log/slog"
	"net/http"

	"github.com/burakkuru5534/case-study-dining2/internal/apperr"
)

const maxBodyBytes = 1 << 20 // 1 MiB

// WriteJSON writes v as JSON with the given status.
func WriteJSON(w http.ResponseWriter, status int, v any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if v == nil {
		return
	}
	if err := json.NewEncoder(w).Encode(v); err != nil {
		slog.Error("write json", "err", err)
	}
}

// ReadJSON decodes a single JSON object from the request body into dst.
// Unknown fields and trailing data are rejected.
func ReadJSON(w http.ResponseWriter, r *http.Request, dst any) error {
	r.Body = http.MaxBytesReader(w, r.Body, maxBodyBytes)
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	if err := dec.Decode(dst); err != nil {
		return fmt.Errorf("%w: invalid json: %v", apperr.ErrValidation, err)
	}
	if err := dec.Decode(&struct{}{}); !errors.Is(err, io.EOF) {
		return fmt.Errorf("%w: body must contain a single json object", apperr.ErrValidation)
	}
	return nil
}

type errorBody struct {
	Error string `json:"error"`
}

// WriteError maps err to a status code and writes a JSON error body.
// Internal errors are logged and hidden from the client.
func WriteError(w http.ResponseWriter, r *http.Request, err error) {
	status := apperr.HTTPStatus(err)
	msg := err.Error()
	if status == http.StatusInternalServerError {
		slog.ErrorContext(r.Context(), "internal error", "err", err, "path", r.URL.Path)
		msg = "internal server error"
	}
	WriteJSON(w, status, errorBody{Error: msg})
}
