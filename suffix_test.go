package main

import (
	"testing"
)

func TestGetNextAlphabeticSuffix(t *testing.T) {
	tests := []struct {
		name    string
		suffix  string
		length  int
		want    string
		wantErr bool
	}{
		{
			name:    "empty string, length 3",
			suffix:  "",
			length:  3,
			want:    "xaa",
			wantErr: false,
		},
		{
			name:    "empty string, length 4",
			suffix:  "",
			length:  4,
			want:    "xaaa",
			wantErr: false,
		},
		{
			name:    "increment last character",
			suffix:  "xab",
			length:  3,
			want:    "xac",
			wantErr: false,
		},
		{
			name:    "roll over last character",
			suffix:  "xaz",
			length:  3,
			want:    "xba",
			wantErr: false,
		},
		{
			name:    "roll over all characters",
			suffix:  "xzzz",
			length:  4,
			want:    "",
			wantErr: true,
		},
		{
			name:    "mismatched length",
			suffix:  "xaz",
			length:  2,
			want:    "",
			wantErr: true,
		},
		{
			name:    "invalid character",
			suffix:  "xaz!",
			length:  4,
			want:    "",
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getNextAlphabeticSuffix("x", tt.suffix, tt.length)
			if (err != nil) != tt.wantErr {
				t.Errorf("getNextAlphabeticSuffix() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getNextAlphabeticSuffix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetNextNumericSuffix(t *testing.T) {
	tests := []struct {
		name      string
		suffix    string
		length    int
		expected  string
		expectErr bool
	}{
		{
			name:      "empty suffix",
			suffix:    "",
			length:    3,
			expected:  "x00",
			expectErr: false,
		},
		{
			name:      "increase digit",
			suffix:    "x09",
			length:    3,
			expected:  "x10",
			expectErr: false,
		},
		{
			name:      "rollover to x00",
			suffix:    "x99",
			length:    3,
			expected:  "",
			expectErr: true,
		},
		{
			name:      "illegal suffix length",
			suffix:    "x9",
			length:    3,
			expected:  "",
			expectErr: true,
		},
		{
			name:      "illegal character in suffix",
			suffix:    "xA9",
			length:    3,
			expected:  "",
			expectErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := getNextNumericSuffix("x", tt.suffix, tt.length)
			if err != nil && !tt.expectErr {
				t.Fatalf("unexpected error: %v", err)
			}
			if err == nil && tt.expectErr {
				t.Fatal("expected an error, but got none")
			}
			if result != tt.expected {
				t.Fatalf("expected %q, but got %q", tt.expected, result)
			}
		})
	}
}
