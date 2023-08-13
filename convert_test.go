package main

import (
	"testing"
)

func TestConvertByteSize(t *testing.T) {
	tests := []struct {
		name    string
		input   string
		want    int
		wantErr bool
	}{
		{
			name:    "plain number without suffix",
			input:   "123",
			want:    123,
			wantErr: false,
		},
		{
			name:    "number with K suffix",
			input:   "1K",
			want:    1024,
			wantErr: false,
		},
		{
			name:    "number with lowercase k suffix",
			input:   "1k",
			want:    1024,
			wantErr: false,
		},
		{
			name:    "number with M suffix",
			input:   "2M",
			want:    2 * 1024 * 1024,
			wantErr: false,
		},
		{
			name:    "number with G suffix",
			input:   "3G",
			want:    3 * 1024 * 1024 * 1024,
			wantErr: false,
		},
		{
			name:    "invalid number",
			input:   "1A2K",
			want:    0,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := convertByteSize(tt.input)

			if (err != nil) != tt.wantErr {
				t.Errorf("convertByteSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if got != tt.want {
				t.Errorf("convertByteSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
