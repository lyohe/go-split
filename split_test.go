package main

import (
	"fmt"
	"reflect"
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

func TestBytesSplit(t *testing.T) {
	tests := []struct {
		name    string
		bytes   []byte
		n       int
		want    []string
		wantErr bool
	}{
		{
			name:    "empty bytes",
			bytes:   nil,
			n:       2,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "illegal byte count",
			bytes:   []byte("test"),
			n:       0,
			want:    nil,
			wantErr: true,
		},
		{
			name:    "split into single group",
			bytes:   []byte("test"),
			n:       10,
			want:    []string{"test"},
			wantErr: false,
		},
		{
			name:    "split multiple bytes",
			bytes:   []byte("testtest"),
			n:       4,
			want:    []string{"test", "test"},
			wantErr: false,
		},
		{
			name:    "split with remaining bytes",
			bytes:   []byte("testtes"),
			n:       4,
			want:    []string{"test", "tes"},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := bytesSplit(tt.bytes, tt.n)

			if (err != nil) != tt.wantErr {
				t.Errorf("bytesSplit() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("bytesSplit() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChunksSplit(t *testing.T) {
	tests := []struct {
		input  []byte
		n      int
		output []string
		err    error
	}{
		{
			input:  []byte("1234567890"),
			n:      1,
			output: []string{"1234567890"},
		},
		{
			input:  []byte("1234567890"),
			n:      2,
			output: []string{"12345", "67890"},
		},
		{
			input:  []byte("1234567890"),
			n:      3,
			output: []string{"1234", "5678", "90"},
		},
		{
			input:  []byte("1234567890"),
			n:      10,
			output: []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"},
		},
		{
			input: []byte("1234567890"),
			n:     11,
			err:   fmt.Errorf("can't split into more than 10 files"), // Expected error
		},
		{
			input: []byte("1234567890"),
			n:     0,
			err:   fmt.Errorf("illegal chunk count"), // Expected error
		},
		{
			input: nil,
			n:     2,
			err:   fmt.Errorf("empty text"), // Expected error
		},
	}

	for _, tt := range tests {
		result, err := chunksSplit(tt.input, tt.n)
		if !reflect.DeepEqual(result, tt.output) || (err != nil && tt.err != nil && err.Error() != tt.err.Error()) {
			t.Errorf("expected %v and %v; got %v and %v", tt.output, tt.err, result, err)
		}
	}
}
