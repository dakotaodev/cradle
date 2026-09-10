package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := NewRouter()

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, 200, w.Code)
	assert.JSONEq(t, w.Body.String(), `{"status":"ok"}`)
}
