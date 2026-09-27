package diaper

import (
	"context"
	"testing"
	"time"

	"github.com/dakotaodev/cradle/internal/db"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
)

// For a small test, use a fake db.Querier that records whether CreateDiaperEvent was called.
//  Check that malformed input returns an error without calling it. Then run go test ./....

type FakeQuerier struct {
	createCalled bool
}

func (f *FakeQuerier) CreateDiaperEvent(ctx context.Context, arg db.CreateDiaperEventParams) (db.DiaperEvent, error) {
	f.createCalled = true
	return db.DiaperEvent{
		ID:         pgtype.UUID{},
		BabyID:     arg.BabyID,
		DiaperType: arg.DiaperType,
		Notes:      arg.Notes,
		OccurredAt: arg.OccurredAt,
		CreatedAt:  pgtype.Timestamptz{},
	}, nil
}

func TestFakeRepository(t *testing.T) {

	testCases := []struct {
		name    string
		input   CreateInput
		wantErr bool
	}{
		{
			name: "valid input",
			input: CreateInput{
				BabyID:     uuid.NewString(),
				Type:       TypeWet,
				OccurredAt: time.Now(),
				Notes:      "",
			},
			wantErr: false,
		},
		{
			name: "invalid uuid",
			input: CreateInput{
				BabyID:     "not-uuid",
				Type:       TypeWet,
				OccurredAt: time.Now(),
				Notes:      "",
			},
			wantErr: true,
		},
		{
			name: "invalid type",
			input: CreateInput{
				BabyID:     uuid.NewString(),
				Type:       "loaded",
				OccurredAt: time.Now(),
				Notes:      "",
			},
			wantErr: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			fake := FakeQuerier{}
			repo := NewRepository(&fake)
			event, err := repo.Create(t.Context(), tc.input)
			if tc.wantErr {
				if fake.createCalled == true {
					t.Error("create was called on malformed input. created should not have been reached.")
				}
				if err == nil {
					t.Error("error did not occur for invalid input.")
				}
			}
			if !tc.wantErr {
				if fake.createCalled != true {
					t.Error("create was not called on valid input")
				}
				if tc.input.BabyID != event.BabyID {
					t.Errorf("the created event baby IDs do not match")
				}
			}
		})
	}
}
