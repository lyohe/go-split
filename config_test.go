package main

import (
	"os"
	"reflect"
	"testing"
)

func TestParseFlags(t *testing.T) {
	tests := []struct {
		name       string
		inputArgs  []string
		wantConfig *SplitConfig
		wantErr    bool
	}{
		{
			name:      "default values",
			inputArgs: []string{"cmd"},
			wantConfig: &SplitConfig{
				SplitType:    SplitByLines,
				SuffixType:   Alphabetic,
				Prefix:       "x",
				SuffixLength: 3,
				LinesCount:   1000,
			},
			wantErr: false,
		},
		{
			name:      "set lines count",
			inputArgs: []string{"cmd", "-l", "500"},
			wantConfig: &SplitConfig{
				SplitType:    SplitByLines,
				SuffixType:   Alphabetic,
				Prefix:       "x",
				SuffixLength: 3,
				LinesCount:   500,
			},
			wantErr: false,
		},
		{
			name:      "set bytes count",
			inputArgs: []string{"cmd", "-b", "100K"},
			wantConfig: &SplitConfig{
				SplitType:    SplitByBytes,
				SuffixType:   Alphabetic,
				Prefix:       "x",
				SuffixLength: 3,
				LinesCount:   1000,
				BytesCount:   100 * 1024, // 100K in bytes (assuming convertByteSize works this way)
			},
			wantErr: false,
		},
		{
			name:      "use numeric suffix",
			inputArgs: []string{"cmd", "-d"},
			wantConfig: &SplitConfig{
				SplitType:    SplitByLines,
				SuffixType:   Numeric,
				Prefix:       "x",
				SuffixLength: 3,
				LinesCount:   1000,
			},
			wantErr: false,
		},
		{
			name:       "too many arguments",
			inputArgs:  []string{"cmd", "file1", "prefix", "extraArg"},
			wantConfig: nil,
			wantErr:    true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Mock the command-line arguments
			os.Args = tt.inputArgs

			gotConfig, _, err := ParseFlags()

			if (err != nil) != tt.wantErr {
				t.Errorf("ParseFlags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(gotConfig, tt.wantConfig) {
				t.Errorf("ParseFlags() = %v, want %v", gotConfig, tt.wantConfig)
			}
		})
	}
}
