package gorandom

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	alphanumeric = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	alpha        = "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
	numeric      = "0123456789"
	distinct     = "2345679ACDEFHJKLMNPRSTUVWXYZ"
	hexdec       = "0123456789abcdef"
	sex          = "fm"
)

var (
	femaleNamesCount *big.Int
	maleNamesCount   *big.Int
	surnamesCount    *big.Int
)

func init() {
	femaleNamesCount = big.NewInt(int64(len(femaleNames)))
	maleNamesCount = big.NewInt(int64(len(maleNames)))
	surnamesCount = big.NewInt(int64(len(surnames)))
}

func AlphaNumeric(n int) string {
	return Random(n, alphanumeric)
}

func Alpha(n int) string {
	return Random(n, alpha)
}

func Numeric(n int) string {
	return Random(n, numeric)
}

func Sex() string {
	return Random(1, sex)
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

func MaleName() string {
	rF, _ := rand.Int(rand.Reader, maleNamesCount)
	rS, _ := rand.Int(rand.Reader, surnamesCount)
	return maleNames[rF.Int64()] + " " + surnames[rS.Int64()]
}

func FemaleName() string {
	rF, _ := rand.Int(rand.Reader, femaleNamesCount)
	rS, _ := rand.Int(rand.Reader, surnamesCount)
	return femaleNames[rF.Int64()] + " " + surnames[rS.Int64()]
}

func FullName() string {
	// Are we doing male or female?
	if Sex() == "m" {
		return MaleName()
	} else {
		return FemaleName()
	}
}

func FullNames(n int) []string {
	s := []string{}

	for i := 0; i < n; i++ {
		s = append(s, FullName())
	}

	return s
}

func Split(s string) (string, string) {
	x := strings.Split(s, " ")

	return x[0], x[1]
}
