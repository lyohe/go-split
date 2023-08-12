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
			got, err := getNextAlphabeticSuffix(tt.suffix, tt.length)
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
