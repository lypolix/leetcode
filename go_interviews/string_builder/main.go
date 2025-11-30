package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.Grow(128)
	parts := []string{"Hello", ", ", "world", "!"}
	for _, part := range parts {
		b.WriteString(part)
	}
	s := b.String()
	fmt.Println(s)
}
