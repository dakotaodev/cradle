package api

import (
	"context"
	"log"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/dakotaodev/cradle/internal/diaper"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type FakeRepository struct{}

func (f FakeRepository) Create(ctx context.Context, input diaper.CreateInput) (diaper.Event, error) {
	log.Printf("Received the following input: %v", input)
	err := input.Validate()
	if err != nil {
		return diaper.Event{}, err
	}
	return diaper.Event{
		ID:         "test-id",
		BabyID:     "test-baby-id",
		OccurredAt: time.Now(),
		CreatedAt:  time.Now(),
		Type:       "wet",
		Notes:      "it's a trap!",
	}, nil
}

func TestRouter(t *testing.T) {
	gin.SetMode(gin.TestMode)
	repo := FakeRepository{}
	router := NewRouter(repo)

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

	request = httptest.NewRequest(
		http.MethodPost,
		"/babies/550e8400-e29b-41d4-a716-446655440000/diapers",
		strings.NewReader(`	{
			"notes": "this is a note",
			"occurredAt": "2026-09-26T12:00:00Z",
			"diaperType": "wet"
		}`),
	)
	request.Header.Set("Content-Type", "application/json")
	w = httptest.NewRecorder()
	router.ServeHTTP(w, request)

	assert.Equal(t, http.StatusCreated, w.Code)
}
