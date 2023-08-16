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

const (
	DEFAULT_SPLIT_TYPE    = SplitByLines
	DEFAULT_PREFIX        = "x"
	DEFAULT_SUFFIX_LENGTH = 3
	DEFAULT_LINES_COUNT   = 1000
	DEFAULT_BYTES_COUNT   = ""
	DEFAULT_CHUNKS_COUNT  = 0
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

	suffixLength := fs.Int("a", DEFAULT_SUFFIX_LENGTH, "")
	linesCount := fs.Int("l", DEFAULT_LINES_COUNT, "")
	bytesCount := fs.String("b", DEFAULT_BYTES_COUNT, "")
	chunksCount := fs.Int("n", DEFAULT_CHUNKS_COUNT, "")
	useNumericSuffix := fs.Bool("d", false, "")

	err := fs.Parse(os.Args[1:])
	if err != nil {
		return nil, nil, err
	}
	if fs.NArg() > 2 {
		return nil, nil, fmt.Errorf("too many arguments")
	}

	setFlags := make(map[string]bool)
	fs.Visit(func(f *flag.Flag) { setFlags[f.Name] = true })
	if setFlags["l"] && setFlags["b"] || setFlags["l"] && setFlags["n"] || setFlags["b"] && setFlags["n"] {
		fs.Usage()
		return nil, nil, fmt.Errorf("multiple split options selected")
	}

	var prefix string
	if flag.NArg() == 2 {
		prefix = fs.Arg(1)
	} else {
		prefix = DEFAULT_PREFIX
	}

	config := &SplitConfig{
		Prefix:       prefix,
		SplitType:    DEFAULT_SPLIT_TYPE,
		SuffixLength: *suffixLength,
		LinesCount:   *linesCount,
	}

	if setFlags["b"] {
		config.SplitType = SplitByBytes

		convertedBytesCount, convertErr := convertByteSize(*bytesCount)
		if convertErr != nil {
			return nil, nil, convertErr
		}
		config.BytesCount = convertedBytesCount
	}
	if setFlags["n"] {
		config.SplitType = SplitByChunks
		config.ChunksCount = *chunksCount
	}

	if *useNumericSuffix {
		config.SuffixType = Numeric
	}

	return config, fs, nil
}
