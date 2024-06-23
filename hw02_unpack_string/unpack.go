package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode/utf8"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(input string) (string, error) {
	var builder strings.Builder
	var last string
	size := utf8.RuneCountInString(input)

	for i, r := range []rune(input) { // not `range input` for correct 'i' rune counter
		char := string(r)

		d, errCD := strconv.Atoi(char) // is current digit
		_, errLD := strconv.Atoi(last) // is last digit

		if errCD == nil && (errLD == nil || i == 0) { // two digits in row or first is digit
			return "", ErrInvalidString
		}

		if errCD != nil { // rune
			if len(last) > 0 && errLD != nil {
				builder.WriteString(last)
			}
			if i == size-1 {
				builder.WriteRune(r)
			}
		} else { // digit
			builder.WriteString(strings.Repeat(last, d))
		}

		last = char
	}

	return builder.String(), nil
}
