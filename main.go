package main

import (
	"fmt"
	"io"
	"os"
	"sync"
)

func main() {
	config, fs, parseErr := ParseFlags()
	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "split: %s\n", parseErr.Error())
		os.Exit(2)
	}

	var input []byte
	var inputErr error

	// 引数でファイル入力を受け取る
	if fs.NArg() >= 1 {
		file, inputErr := os.Open(fs.Arg(0))
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
	if fs.NArg() < 1 {
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
	case SplitByChunks:
		splitText, splitErr = chunksSplit(input, config.ChunksCount)
	}
	if splitErr != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", splitErr)
		os.Exit(1)
	}

	// 分割した文字列をファイルに書き込む
	suffix := ""
	var suffixErr error

	var wg sync.WaitGroup
	for _, s := range splitText {
		suffix, suffixErr = getNextSuffix(suffix, config)
		if suffixErr != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", suffixErr)
			os.Exit(1)
		}

		wg.Add(1)
		go func(text, suffix string) {
			defer wg.Done()

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

			_, err = f.WriteString(text)
			if err != nil {
				fmt.Fprintf(os.Stderr, "split: %v\n", err)
				os.Exit(1)
			}
		}(s, suffix)
	}
	wg.Wait()

	os.Exit(0)
}
