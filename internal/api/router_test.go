package api

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := NewRouter()

	w := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/health", nil)

	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, `{"status": "ok"}`, w.Body.String())

	request = httptest.NewRequest(
		http.MethodPost,
		"/babies/550e8400-e29b-41d4-a716-446655440000/diapers",
		strings.NewReader(`	{
			"notes": "this is a note",
			"occurredAt": "2026-09-26T12:00:00Z",
			"diaperType": "wet""
		}`),
	)
	request.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusBadRequest, w.Code)

		request = httptest.NewRequest(
		http.MethodPost,
		"/babies/not-uuid/diapers",
		strings.NewReader(`	{
			"notes": "this is a note",
			"occurredAt": "2026-09-26T12:00:00Z",
			"diaperType": "wet"
		}`),
	)
	request.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusBadRequest, w.Code)
}
