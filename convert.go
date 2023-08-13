package main

import (
	"strconv"
)

func convertByteSize(sizeStr string) (int, error) {
	lastChar := sizeStr[len(sizeStr)-1]

	var size string
	var uint byte

	if '0' <= lastChar && lastChar <= '9' {
		size = sizeStr
	} else {
		size = sizeStr[:len(sizeStr)-1]
		uint = lastChar
	}

	value, err := strconv.Atoi(size)
	if err != nil {
		return 0, err
	}

	switch uint {
	case 'K', 'k':
		return value * 1024, nil
	case 'M', 'm':
		return value * 1024 * 1024, nil
	case 'G', 'g':
		return value * 1024 * 1024 * 1024, nil
	default:
		return value, nil
	}
}
