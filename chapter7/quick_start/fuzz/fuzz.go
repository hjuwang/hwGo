package fuzz

import (
	"errors"
	"unicode/utf8"
)

func ReverseByBytes(s string) (string, error) {

	if !utf8.ValidString(s) {
		return s, errors.New("invalid string")
	}

	b := []byte(s)

	for i, j := 0, len(b)-1; i < len(b)/2; i, j = i+1, j-1 {
		b[i], b[j] = b[j], b[i]
	}

	return string(b), nil
}

func ReverseByRunes(s string) (string, error) {

	if !utf8.ValidString(s) {
		return s, errors.New("invalid string")
	}

	r := []rune(s)

	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
		r[i], r[j] = r[j], r[i]
	}

	return "", nil
}
