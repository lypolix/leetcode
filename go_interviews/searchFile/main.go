package main

import (
	"os"
	"regexp"
)

func main() {

}

var digitRegexp = regexp.MustCompile("[0-9]+")

func FindDigits(filename string) []byte {
	b, _ := os.ReadFile(filename)

	b = digitRegexp.Find(b)

	res := make([]byte, len(b))

	copy(res, b)

	return res
}
