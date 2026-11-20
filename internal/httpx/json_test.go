package httpx

import (
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/burakkuru5534/case-study-dining2/internal/apperr"
)

func TestReadJSON(t *testing.T) {
	type payload struct {
		Name string `json:"name"`
	}

	tests := []struct {
		name    string
		body    string
		wantErr bool
	}{
		{"valid", `{"name":"pizza"}`, false},
		{"unknown field", `{"name":"pizza","x":1}`, true},
		{"malformed", `{"name":`, true},
		{"trailing data", `{"name":"a"}{"name":"b"}`, true},
		{"empty body", ``, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(tt.body))
			w := httptest.NewRecorder()

			var p payload
			err := ReadJSON(w, r, &p)
			if (err != nil) != tt.wantErr {
				t.Fatalf("err = %v, wantErr %v", err, tt.wantErr)
			}
			if err != nil && !errors.Is(err, apperr.ErrValidation) {
				t.Fatalf("expected ErrValidation, got %v", err)
			}
		})
	}
}

func TestWriteError_HidesInternalErrors(t *testing.T) {
	r := httptest.NewRequest(http.MethodGet, "/", nil)
	w := httptest.NewRecorder()

	WriteError(w, r, errors.New("db password is hunter2"))

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("status = %d", w.Code)
	}
	if strings.Contains(w.Body.String(), "hunter2") {
		t.Fatalf("internal error leaked: %s", w.Body.String())
	}
}
