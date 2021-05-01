package utils

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

const (
	/*
		base62Range is the possible different characters in short url unique identifier
		let's say our short url is: https://www.localhost:8080/ab53hRdpZhf
		then ab53hRdpZhf is the unique identifier
	*/
	base62Range = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

	length = uint64(len(base62Range))
)

// ConvertUniqueKeyToUrlPath : does the base62 encoding of the unique key
func ConvertUniqueKeyToUrlPath(uniqueKey uint64) string {
	var encodedString string
	for uniqueKey > 0 {
		encodedString += string(base62Range[(uniqueKey % length)])
		uniqueKey = uniqueKey / length
	}
	return encodedString
}

func ConvertUrlPathToUniqueKey(urlPath string) (uint64, error) {
	var uniqueKey uint64

	for i, symbol := range urlPath {
		alphabeticPosition := strings.IndexRune(base62Range, symbol)
		// if the symbol is not present in the base62Range
		if alphabeticPosition == -1 {
			invalidCharError := fmt.Sprintf("Invalid character present in the url path: '%s'", string(symbol))
			return uint64(0), errors.New(invalidCharError)
		}
		uniqueKey += uint64(alphabeticPosition) * uint64(math.Pow(float64(length), float64(i)))
	}

	return uniqueKey, nil
}
