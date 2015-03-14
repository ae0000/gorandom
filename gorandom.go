package gorandom

import "crypto/rand"

const (
	alphanumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	alpha        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numeric      = "0123456789"
	distinct     = "2345679ACDEFHJKLMNPRSTUVWXYZ"
	hexdec       = "0123456789abcdef"
)

func AlphaNumeric(n int) string {
	return Random(n, alphanumeric)
}

func Alpha(n int) string {
	return Random(n, alpha)
}

func Numeric(n int) string {
	return Random(n, numeric)
}

func Distinct(n int) string {
	return Random(n, distinct)
}

func Hexdec(n int) string {
	return Random(n, hexdec)
}

func Random(n int, pool string) string {
	var bytes = make([]byte, n)
	rand.Read(bytes)
	for i, b := range bytes {
		bytes[i] = pool[b%byte(len(pool))]
	}
	return string(bytes)
}
