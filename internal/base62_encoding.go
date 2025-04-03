package internal

import (
	"strconv"
)

func getAsciiValue(index int) string {
	var asciiValue rune 

	if index <= 9 {
		return strconv.Itoa(index)
	} else if index <= 35 {
		asciiValue = rune(97 + (index - 10))
	} else if index <= 61 {
		asciiValue = rune(65 + (index - 36))
	}
	return string(asciiValue)
}

func GetBase62Encoding(value int64) string {
	var encodedValue string

	for value > 0 {
		encodedValue = getAsciiValue(int(value % 62)) + encodedValue
		value /= 62
	}
	return encodedValue
}
