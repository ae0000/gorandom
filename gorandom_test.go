package gorandom

import (
	"fmt"
	"regexp"
	"testing"
)

var length = 100 // Length of characters
func TestAlpha(t *testing.T) {
	a := Alpha(length)

	if len(a) != length {
		t.Error("Length failed!")
	}

	r := regexp.MustCompile("^[a-zA-Z]+$")

	if !r.MatchString(a) {
		t.Errorf("Non-alphas made it into the alphas??? %s", a)
	}
}

func TestAlphaNumeric(t *testing.T) {
	a := AlphaNumeric(length)

	if len(a) != length {
		t.Error("Length failed!")
	}

	r := regexp.MustCompile("^[a-zA-Z0-9]+$")

	if !r.MatchString(a) {
		t.Errorf("Non-alphanumerics made it in??? %s", a)
	}
}

func TestNumeric(t *testing.T) {
	a := Numeric(length)

	if len(a) != length {
		t.Error("Length failed!")
	}

	r := regexp.MustCompile("^[0-9]+$")

	if !r.MatchString(a) {
		t.Errorf("Non-numerics made it in??? %s", a)
	}
}

func TestHa(t *testing.T) {
	a := FullNames(20)

	fmt.Println(a)
}
