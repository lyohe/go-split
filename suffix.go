package main

import (
	"fmt"
	"strings"
	"unicode"
)

func getNextSuffix(suffix string, config *SplitConfig) (string, error) {
	switch config.SuffixType {
	case Alphabetic:
		return getNextAlphabeticSuffix(config.Prefix, suffix, config.SuffixLength)
	case Numeric:
		return getNextNumericSuffix(config.Prefix, suffix, config.SuffixLength)
	}
	return "", fmt.Errorf("illegal suffix type")
}

func getNextAlphabeticSuffix(prefix string, suffix string, length int) (string, error) {
	if length < 1 || (len(suffix)-len(prefix)+1 != length && suffix != "") {
		return "", fmt.Errorf("illegal suffix length")
	}

	if suffix == prefix+strings.Repeat("z", length-1) {
		return "", fmt.Errorf("too many files")
	}

	for _, r := range suffix {
		if !unicode.IsLetter(r) {
			return "", fmt.Errorf("illegal character")
		}
	}

	if suffix == "" {
		return prefix + strings.Repeat("a", length-1), nil
	}

	chars := []rune(suffix)
	for i := length - 1; i >= 0; i-- {
		if chars[i] != 'z' {
			chars[i]++
			return string(chars), nil
		}
		chars[i] = 'a'
	}
	return string(chars), nil
}

func getNextNumericSuffix(prefix string, suffix string, length int) (string, error) {
	if length < 1 || (len(suffix)-len(prefix)+1 != length && suffix != "") {
		return "", fmt.Errorf("illegal suffix length")
	}

	if suffix == prefix+strings.Repeat("9", length-1) {
		return "", fmt.Errorf("too many files")
	}

	for i, r := range suffix {
		if i != 0 && !unicode.IsDigit(r) {
			return "", fmt.Errorf("illegal character")
		}
	}

	if suffix == "" {
		return prefix + strings.Repeat("0", length-1), nil
	}

	chars := []rune(suffix)
	for i := length - 1; i >= 0; i-- {
		if chars[i] != '9' {
			chars[i]++
			return string(chars), nil
		}
		chars[i] = '0'
	}
	return string(chars), nil
}
