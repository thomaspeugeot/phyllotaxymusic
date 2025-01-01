package tests

import (
	"testing"

	"github.com/thomaspeugeot/phyllotaxymusic/go/models"
)

func TestToggleNotePlayed(t *testing.T) {
	tests := []struct {
		name                  string
		initialEncoding       int
		toggles               []int // multiple toggles in sequence
		expectedFinalEncoding int
	}{
		{
			name:                  "Toggle bit 0 from 0 => 1",
			initialEncoding:       0,
			toggles:               []int{0},
			expectedFinalEncoding: 1, // 0001
		},
		{
			name:                  "Toggle bit 1 from 1 => 3",
			initialEncoding:       1, // 0001
			toggles:               []int{1},
			expectedFinalEncoding: 3, // 0011
		},
		{
			name:                  "Toggle bit 0 twice from 1 => 1",
			initialEncoding:       1, // 0001
			toggles:               []int{0, 0},
			expectedFinalEncoding: 1, // Toggles bit 0 twice => no net change
		},
		{
			name:                  "No-op for out of range (-1)",
			initialEncoding:       5,         // 0101
			toggles:               []int{-1}, // invalid toggle
			expectedFinalEncoding: 5,         // No change
		},
		{
			name:                  "No-op for out of range (64)",
			initialEncoding:       5,         // 0101
			toggles:               []int{64}, // invalid toggle
			expectedFinalEncoding: 5,         // No change
		},
		{
			name:                  "Toggle bits 0,1,2 on an empty encoding => 7",
			initialEncoding:       0,
			toggles:               []int{0, 1, 2},
			expectedFinalEncoding: 7, // 0111
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			parameter := &models.Parameter{
				ThemeBinaryEncoding: tt.initialEncoding,
			}

			// Perform each toggle in sequence
			for _, toggle := range tt.toggles {
				parameter.ToggleNotePlayed(toggle)
			}

			if parameter.ThemeBinaryEncoding != tt.expectedFinalEncoding {
				t.Errorf(
					"ToggleNotePlayed() => got %d, want %d",
					parameter.ThemeBinaryEncoding, tt.expectedFinalEncoding,
				)
			}
		})
	}
}
