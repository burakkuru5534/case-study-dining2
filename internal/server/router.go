// Package server wires handlers and middleware into an http.Handler.
package server

import (
	"net/http"

	"github.com/burakkuru5534/case-study-dining2/internal/httpx"
)

// NewRouter builds the application's HTTP handler.
// Register new routes here using Go 1.22+ patterns, e.g. "POST /orders", "GET /orders/{id}".
func NewRouter() http.Handler {
	mux := http.NewServeMux()

	mux.HandleFunc("GET /healthz", func(w http.ResponseWriter, r *http.Request) {
		httpx.WriteJSON(w, http.StatusOK, map[string]string{"status": "ok"})
	})

	return httpx.Chain(mux,
		httpx.WithRecover,
		httpx.WithRequestID,
		httpx.WithLogging,
	)
}
