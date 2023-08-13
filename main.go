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
)

const (
	Alphabetic SuffixType = iota
	Numeric
)

type SplitConfig struct {
	SplitType    SplitType
	SuffixType   SuffixType
	SuffixLength int
	LinesCount   int
	BytesCount   int
}

func ParseFlags() *SplitConfig {
	var parseErr error

	suffixLength := flag.Int("a", 3, "use suffixes of length N (default 3)")
	linesCount := flag.Int("l", 1000, "put N lines/records per output file")
	bytesCount := flag.String("b", "", "Specify bytes per file with optional multiplier (k, K, m, M, g, G)")
	useNumericSuffix := flag.Bool("d", false, "use numeric suffixes instead of alphabetic")
	flag.Parse()

	config := &SplitConfig{
		LinesCount:   *linesCount,
		SuffixLength: *suffixLength,
	}
	if *bytesCount != "" {
		config.SplitType = SplitByBytes
		config.BytesCount, parseErr = convertByteSize(*bytesCount)
	} else {
		config.SplitType = SplitByLines
	}

	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", parseErr)
		flag.Usage()
		os.Exit(2)
	}

	if *useNumericSuffix {
		config.SuffixType = Numeric
	}
	/*
		// 引数が正しくない場合は usage を表示して終了
		if flag.NArg() > 1 {
			fmt.Fprintf(os.Stderr, "split: too many arguments\n")
			flag.Usage()
			os.Exit(2)
		}
	*/

	return config
}

func init() {
	flag.Usage = usage
}

func main() {
	config := ParseFlags()

	var input []byte
	var inputErr error

	// 引数でファイル入力を受け取る
	if flag.NArg() == 1 {
		file, inputErr := os.Open(flag.Arg(0))
		if inputErr != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", inputErr)
			os.Exit(1)
		}
		defer file.Close()

		input, inputErr = io.ReadAll(file)
		if inputErr != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", inputErr)
			os.Exit(1)
		}
	}

	// ファイル入力がない場合は標準入力を受け取る
	if flag.NArg() < 1 {
		input, inputErr = io.ReadAll(os.Stdin)
		if inputErr != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", inputErr)
			os.Exit(1)
		}
	}

	var splitText []string
	var splitErr error

	// 入力を []string として分割する
	switch config.SplitType {
	case SplitByLines:
		splitText, splitErr = linesSplit(string(input), config.LinesCount)
	case SplitByBytes:
		splitText, splitErr = bytesSplit(input, config.BytesCount)
	}
	if splitErr != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", splitErr)
		os.Exit(1)
	}

	// 分割した文字列をファイルに書き込む
	suffix := ""
	var suffixErr error
	for _, s := range splitText {
		suffix, suffixErr = getNextSuffix(suffix, config)
		if suffixErr != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", suffixErr)
			os.Exit(1)
		}
		f, err := os.Create(suffix)
		if err != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", err)
			os.Exit(1)
		}
		defer func() {
			if err := f.Close(); err != nil {
				fmt.Fprintf(os.Stderr, "split: %v\n", err)
				os.Exit(1)
			}
			f.Close()
		}()

		_, err = f.WriteString(s)
		if err != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", err)
			os.Exit(1)
		}
	}

	os.Exit(0)
}
