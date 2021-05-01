package utils

const (
	/*
		base62Range is the possible different characters in short url unique identifier
		let's say our short url is: https://www.localhost:8080/ab53hodpdhf
		then ab53hodpdhf is the unique identifier
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
