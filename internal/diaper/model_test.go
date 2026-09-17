package diaper

import (
	"strings"
	"testing"
	"time"
)

func TestDiaperValidation(t *testing.T) {
	tests := []struct {
		name    string
		input   Diaper
		wantErr bool
	}{
		{
			name: "empty BabyID",
			input: Diaper{
				BabyID:     "",
				OccurredAt: time.Now(),
				Type:       TypeDry,
				Notes:      "test",
			},
			wantErr: true,
		},
		{
			name: "Notes too long",
			input: Diaper{
				BabyID:     "m",
				OccurredAt: time.Now(),
				Type:       TypeWet,
				Notes:      strings.Repeat("a", 501),
			},
			wantErr: true,
		},
		{
			name: "Invalid diaper type",
			input: Diaper{
				BabyID:     "d",
				OccurredAt: time.Now(),
				Type:       "moist",
				Notes:      "test",
			},
			wantErr: true,
		},
		{
			name: "Occurred in future",
			input: Diaper{
				BabyID:     "d",
				OccurredAt: time.Now().Add(time.Hour),
				Type:       TypeDry,
				Notes:      "no notes",
			},
			wantErr: true,
		},
		{
			name: "Occurred in zero time",
			input: Diaper{
				BabyID:     "d",
				OccurredAt: time.Time{},
				Type:       TypeDry,
				Notes:      "no notes",
			},
			wantErr: true,
		},
		{
			name: "valid diaper",
			input: Diaper{
				BabyID:     "ID",
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
