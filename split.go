package main

import (
	"fmt"
	"strings"
)

func linesSplit(text string, n int) ([]string, error) {
	if n < 1 {
		return nil, fmt.Errorf("illegal line count") // TODO: error message
	}

	if text == "" {
		return nil, fmt.Errorf("empty text") // TODO: error message
	}

	splitLines := strings.Split(text, "\n")
	length := len(splitLines)
	if length <= n {
		return []string{text + "\n"}, nil
	}

	var lines []string
	line := ""
	count := 0
	for _, s := range splitLines {
		if count < n {
			line = line + s + "\n"
			count++
		}
		if count == n || s == splitLines[length-1] {
			lines = append(lines, line)
			line = ""
			count = 0
		}
	}

	return lines, nil
}

func bytesSplit(bytes []byte, n int) ([]string, error) {
	if n < 1 {
		return nil, fmt.Errorf("illegal byte count") // TODO: error message
	}

	if bytes == nil {
		return nil, fmt.Errorf("empty bytes") // TODO: error message
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
