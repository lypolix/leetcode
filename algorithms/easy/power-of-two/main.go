package main

import "fmt"

func main() {
	fmt.Println(isPowerOfTwo(128))
}

func isPowerOfTwo(n int) bool {
    i := 1
    for i <= n {
        if i == n {
            return true
        }
        i *= 2
    }
    return false
}
