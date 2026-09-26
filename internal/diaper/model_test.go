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
				BabyID:     "m",
				OccurredAt: time.Now(),
				Type:       TypeWet,
				Notes:      strings.Repeat("a", 501),
			},
			wantErr: true,
		},
		{
			name: "Invalid diaper type",
			input: CreateInput{
				BabyID:     "d",
				OccurredAt: time.Now(),
				Type:       "moist",
				Notes:      "test",
			},
			wantErr: true,
		},
		{
			name: "Occurred in future",
			input: CreateInput{
				BabyID:     "d",
				OccurredAt: time.Now().Add(time.Hour),
				Type:       TypeDry,
				Notes:      "no notes",
			},
			wantErr: true,
		},
		{
			name: "Occurred in zero time",
			input: CreateInput{
				BabyID:     "d",
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
