package api

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
	
	"github.com/stretchr/testify/assert"
	"github.com/dakotaodev/cradle/internal/babies"
	"github.com/gin-gonic/gin"
)

type fakeBabyRepository struct {
	baby  babies.Baby
	err   error
	gotID string
	calls int
}

func (f *fakeBabyRepository) CreateBaby(context.Context, babies.CreateInput) (babies.Baby, error) {
	panic("unexpected CreateBaby call")
}

func (f *fakeBabyRepository) GetBaby(_ context.Context, id string) (babies.Baby, error) {
	f.calls++
	f.gotID = id
	return f.baby, f.err
}

func TestGetBabyHandler(t *testing.T) {
	gin.SetMode(gin.TestMode)
	baby := babies.Baby{
		ID:        "550e8400-e29b-41d4-a716-446655440000",
		FullName:  "Test Baby",
		BirthDate: time.Date(2026, 6, 25, 0, 0, 0, 0, time.UTC),
		CreatedAt: time.Date(2026, 10, 20, 12, 0, 0, 0, time.UTC),
	}
	tests := []struct {
		name       string
		err        error
		wantStatus int
	}{
		{name: "existing baby", wantStatus: http.StatusOK},
		{name: "missing baby", err: babies.ErrNotFound, wantStatus: http.StatusNotFound},
		{name: "wrapped not found", err: fmt.Errorf("lookup baby: %w", babies.ErrNotFound), wantStatus: http.StatusNotFound},
		{name: "database failure", err: errors.New("database unavailable"), wantStatus: http.StatusInternalServerError},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repo := &fakeBabyRepository{baby: baby, err: tt.err}
			router := NewRouter(nil, repo)
			rec := httptest.NewRecorder()
			req := httptest.NewRequest(http.MethodGet, "/babies/"+baby.ID, nil)
			router.ServeHTTP(rec, req)
			if rec.Code != tt.wantStatus {
				t.Fatalf("status = %d, want %d; body: %s", rec.Code, tt.wantStatus, rec.Body.String())
			}
			if repo.calls != 1 || repo.gotID != baby.ID {
				t.Fatalf("GetBaby calls = %d, ID = %q; want 1, %q", repo.calls, repo.gotID, baby.ID)
			}
			if tt.err != nil {
				return
			}
			assert.JSONEq(t, `{
    		"id": "550e8400-e29b-41d4-a716-446655440000",
    		"fullName": "Test Baby",
    		"birthDate": "2026-06-25",
    		"createdAt": "2026-10-20T12:00:00Z"
				}`, rec.Body.String())
			})
	}
}
