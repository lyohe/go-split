package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

type SplitType int
type SuffixType int

const (
	SplitByLines SplitType = iota
	SplitByBytes
	SplitByChunks
)

const (
	Alphabetic SuffixType = iota
	Numeric
)

type SplitConfig struct {
	SplitType    SplitType
	SuffixType   SuffixType
	Prefix       string
	SuffixLength int
	LinesCount   int
	BytesCount   int
	ChunksCount  int
}

func usage() {
	fmt.Fprintln(os.Stderr, "usage: split [-l line_count] [-a suffix_length] [file [prefix]]\n       split -b byte_count[K|k|M|m|G|g] [-a suffix_length] [file [prefix]]\n       split -n chunk_count [-a suffix_length] [file [prefix]]\n       split -p pattern [-a suffix_length] [file [prefix]]")
}

func ParseFlags() (*SplitConfig, *flag.FlagSet, error) {
	fs := flag.NewFlagSet(os.Args[0], flag.ContinueOnError)

	fs.Usage = usage
	fs.SetOutput(io.Discard)

	suffixLength := fs.Int("a", 3, "")
	linesCount := fs.Int("l", 1000, "")
	bytesCount := fs.String("b", "", "")
	chunksCount := fs.Int("n", 0, "")
	useNumericSuffix := fs.Bool("d", false, "")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return nil, nil, err
	}
	if fs.NArg() > 2 {
		return nil, nil, fmt.Errorf("too many arguments")
	}

	var prefix string
	if flag.NArg() == 2 {
		prefix = fs.Arg(1)
	} else {
		prefix = "x"
	}

	config := &SplitConfig{
		Prefix:       prefix,
		SuffixLength: *suffixLength,
		LinesCount:   *linesCount,
	}

	var parseErr error

	if *bytesCount != "" {
		config.SplitType = SplitByBytes
		config.BytesCount, parseErr = convertByteSize(*bytesCount)
	} else if *chunksCount != 0 {
		config.SplitType = SplitByChunks
		config.ChunksCount = *chunksCount
	} else {
		config.SplitType = SplitByLines
	}

	if parseErr != nil {
		return nil, nil, parseErr
	}

	if *useNumericSuffix {
		config.SuffixType = Numeric
	}

	return config, fs, nil
}
