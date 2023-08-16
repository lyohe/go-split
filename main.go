package main

import (
	"flag"
	"fmt"
	"io"
	"os"
	"sync"
)

func readInput(fs *flag.FlagSet) ([]byte, error) {
	if fs.NArg() >= 1 {
		file, err := os.Open(fs.Arg(0))
		if err != nil {
			return nil, fmt.Errorf("%s: No such file or directory", fs.Arg(0))
		}
		defer file.Close()

		return io.ReadAll(file)
	}

	return io.ReadAll(os.Stdin)
}

func splitInput(input []byte, config *SplitConfig) ([]string, error) {
	switch config.SplitType {
	case SplitByLines:
		return linesSplit(string(input), config.LinesCount)
	case SplitByBytes:
		return bytesSplit(input, config.BytesCount)
	case SplitByChunks:
		return chunksSplit(input, config.ChunksCount)
	default:
		return nil, fmt.Errorf("unknown split type")
	}
}

func writeToFile(splitText []string, config *SplitConfig) {
	var wg sync.WaitGroup
	var suffix string
	var suffixErr error

	for _, s := range splitText {
		suffix, suffixErr = getNextSuffix(suffix, config)
		if suffixErr != nil {
			fmt.Fprintf(os.Stderr, "split: %v\n", suffixErr)
			os.Exit(1)
		}

		wg.Add(1)
		go func(text, currentSuffix string) {
			defer wg.Done()

			file, err := os.Create(currentSuffix)
			if err != nil {
				fmt.Fprintf(os.Stderr, "split: %v\n", err)
				os.Exit(1)
			}
			defer file.Close()

			_, err = file.WriteString(text)
			if err != nil {
				fmt.Fprintf(os.Stderr, "split: %v\n", err)
				os.Exit(1)
			}
		}(s, suffix)
	}
	wg.Wait()
}

func main() {
	config, fs, parseErr := ParseFlags()
	if parseErr != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", parseErr)
		os.Exit(2)
	}

	input, inputErr := readInput(fs)
	if inputErr != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", inputErr)
		os.Exit(1)
	}

	splitText, splitErr := splitInput(input, config)
	if splitErr != nil {
		fmt.Fprintf(os.Stderr, "split: %v\n", splitErr)
		os.Exit(1)
	}

	writeToFile(splitText, config)
}
