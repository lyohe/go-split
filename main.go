package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	suffixLength int
	linesCount   int
)

func init() {
	flag.IntVar(&suffixLength, "a", 3, "use suffixes of length N (default 3)")
	flag.IntVar(&linesCount, "l", 1000, "put N lines/records per output file")
}

func main() {
	flag.Usage = usage
	flag.Parse()

	// 引数が正しくない場合は usage を表示して終了
	if flag.NArg() > 1 {
		fmt.Fprintf(os.Stderr, "split: too many arguments\n")
		flag.Usage()
		os.Exit(2)
	}

	var input []byte
	var err error

	// 引数でファイル入力を受け取る
	if flag.NArg() == 1 {
		file, err := os.Open(flag.Arg(0))
		if err != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", err)
			os.Exit(1)
		}
		defer file.Close()

		input, err = io.ReadAll(file)
		if err != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", err)
			os.Exit(1)
		}
	}

	// ファイル入力がない場合は標準入力を受け取る
	if flag.NArg() < 1 {
		input, err = io.ReadAll(os.Stdin)
		if err != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", err)
			os.Exit(1)
		}
	}

	// 入力を []string として分割する
	splitText, err := linesSplit(string(input), linesCount)
	if err != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", err)
		os.Exit(1)
	}

	// 分割した文字列をファイルに書き込む
	suffix := ""
	for _, s := range splitText {
		suffix, err = getNextAlphabeticSuffix(suffix, suffixLength)
		if err != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", err)
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
