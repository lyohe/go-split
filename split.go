package main

import (
	"fmt"
	"strings"
)

func linesSplit(text string, n int) ([]string, error) {
	if n < 1 {
		return nil, fmt.Errorf("illegal line count")
	}

	if text == "" {
		return nil, fmt.Errorf("empty text")
	}

	splitLines := strings.Split(text, "\n")
	length := len(splitLines)
	if length <= n {
		return []string{text + "\n"}, nil
	}

	var lines []string
	line := ""
	count := 0
	for i, s := range splitLines {
		if count < n {
			line = line + s + "\n"
			count++
		}
		if count == n || i == length-1 {
			lines = append(lines, line)
			line = ""
			count = 0
		}
	}

	return lines, nil
}

func bytesSplit(bytes []byte, n int) ([]string, error) {
	if n < 1 {
		return nil, fmt.Errorf("illegal byte count")
	}

	if bytes == nil {
		return nil, fmt.Errorf("empty bytes")
	}

	length := len(bytes)
	if length <= int(n) {
		return []string{string(bytes)}, nil
	}

	var lines []string
	line := ""
	count := 0
	for i, b := range bytes {
		if count < n {
			line = line + string(b)
			count++
		}
		if count == n || i == length-1 {
			lines = append(lines, line)
			line = ""
			count = 0
		}
	}

	return lines, nil
}

func chunksSplit(bytes []byte, n int) ([]string, error) {
	if n <= 0 {
		return nil, fmt.Errorf("illegal chunk count")
	}

	if n == 1 {
		return []string{string(bytes)}, nil
	}

	if bytes == nil {
		return nil, fmt.Errorf("empty text")
	}

	length := len(bytes)
	if length < n {
		return nil, fmt.Errorf("can't split into more than %d files", length)
	}

	var bytesPerChunk int
	if length%n == 0 {
		bytesPerChunk = length / n
	} else {
		bytesPerChunk = length/n + 1
	}

	var chunks []string
	text := ""
	count := 0
	for i, b := range bytes {
		if count < bytesPerChunk {
			text = text + string(b)
			count++
		}
		if count == bytesPerChunk || i == length-1 {
			chunks = append(chunks, text)
			text = ""
			count = 0
		}
	}

	return chunks, nil
}
