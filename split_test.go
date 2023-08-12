package main

import (
	"testing"
)

func TestLinesSplit(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		n       int
		want    []string
		wantErr bool
	}{
		{
			name:    "empty text",
			text:    "",
			n:       2,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "illegal line count",
			text:    "test\ntext",
			n:       0,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "split into single group",
			text:    "line1\nline2",
			n:       3,
			want:    []string{"line1\nline2\n"},
			wantErr: false,
		},
		{
			name:    "split multiple lines",
			text:    "line1\nline2\nline3\nline4",
			n:       2,
			want:    []string{"line1\nline2\n", "line3\nline4\n"},
			wantErr: false,
		},
		{
			name:    "split with remaining lines",
			text:    "line1\nline2\nline3",
			n:       2,
			want:    []string{"line1\nline2\n", "line3\n"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := linesSplit(tt.text, tt.n)
			if (err != nil) != tt.wantErr {
				t.Errorf("lines_split() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(got) != len(tt.want) {
				t.Errorf("lines_split() length = %v, want %v", len(got), len(tt.want))
				return
			}
			for i, line := range got {
				if line != tt.want[i] {
					t.Errorf("lines_split() got[%d] = %v, want %v", i, line, tt.want[i])
				}
			}
		})
	}
}
