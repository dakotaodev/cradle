package diaper

import (
	"strings"
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestCreateInputValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   CreateInput
		wantErr bool
	}{
		{
			name: "empty BabyID",
			input: CreateInput{
				BabyID:     "",
				OccurredAt: time.Now(),
				Type:       TypeDry,
				Notes:      "test",
			},
			wantErr: true,
		},
		{
			name: "Notes too long",
			input: CreateInput{
				BabyID:     uuid.NewString(),
				OccurredAt: time.Now(),
				Type:       TypeWet,
				Notes:      strings.Repeat("a", 501),
			},
			wantErr: true,
		},
		{
			name: "Invalid diaper type",
			input: CreateInput{
				BabyID:     uuid.NewString(),
				OccurredAt: time.Now(),
				Type:       "moist",
				Notes:      "test",
			},
			wantErr: true,
		},
		{
			name: "Occurred in future",
			input: CreateInput{
				BabyID:     uuid.NewString(),
				OccurredAt: time.Now().Add(time.Hour),
				Type:       TypeDry,
				Notes:      "no notes",
			},
			wantErr: true,
		},
		{
			name: "Occurred in zero time",
			input: CreateInput{
				BabyID:     uuid.New().String(),
				OccurredAt: time.Time{},
				Type:       TypeDry,
				Notes:      "no notes",
			},
			wantErr: true,
		},
		{
			name: "valid diaper",
			input: CreateInput{
				BabyID:     uuid.New().String(),
				OccurredAt: time.Now(),
				Type:       TypeDry,
				Notes:      "none",
			},
			wantErr: false,
		},
	}

	for _, testCase := range tests {
		t.Run(testCase.name, func(t *testing.T) {
			err := testCase.input.Validate()
			if err != nil && !testCase.wantErr {
				t.Errorf("error received in valid test case: %v", err)
			}
			if testCase.wantErr && err == nil {
				t.Errorf("error not returned for invalid test case: %v", err)
			}

		})
	}
}

func TestToCreateParams(t *testing.T) {
	test_cases := []struct {
		name    string
		input   CreateInput
		wantErr bool
	}{
		{
			name: "valid input",
			input: CreateInput{
				BabyID:     uuid.NewString(),
				Type:       "wet",
				OccurredAt: time.Now(),
				Notes:      "this should pass!",
			},
			wantErr: false,
		},
		{
			name: "bad uuid for baby ID",
			input: CreateInput{
				BabyID:     "not-an-uuid",
				Type:       "wet",
				OccurredAt: time.Now(),
				Notes:      "disappointment",
			},
			wantErr: true,
		},
		{
			name: "invalid diaper type",
			input: CreateInput{
				BabyID:     uuid.NewString(),
				Type:       "fully loaded",
				OccurredAt: time.Now(),
				Notes:      "time to size up",
			},
			wantErr: true,
		},
	}

	for _, tc := range test_cases {
		t.Run(tc.name, func(t *testing.T) {
			params, err := toCreateParams(tc.input)
			if tc.wantErr {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				return
			}
			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			parsedBabyId, err := uuid.Parse(tc.input.BabyID)
			if err != nil {
				t.Fatalf("unable to parse BabyID as uuid: %v", err)
			} else {
				if params.BabyID.Bytes != parsedBabyId {
					t.Error("BabyID bytes do not match parsed UUID")
				}
			}

			if params.BabyID.Valid != true || params.OccurredAt.Valid != true {
				t.Error("BabyID and OccurredAt must both have Valid set to true.")
			}
			if tc.input.Notes != params.Notes {
				t.Error("notes field does not match after conversion")
			}
			if tc.input.OccurredAt != params.OccurredAt.Time {
				t.Error("OccurredAt field does not match after conversion")
			}
			if string(tc.input.Type) != params.DiaperType {
				t.Error("diaper type field does not match after conversion")
			}
		})
	}
}
