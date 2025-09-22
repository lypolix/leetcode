package main

import (
	"fmt"
	"strings"
)


func main() {
	fmt.Println(multiply("2", "3"))
	fmt.Println(multiply("123", "456"))
}

func multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	res := make([]int, len(num1) + len(num2))

	for i := len(num1) - 1; i >= 0; i -- {
		for j := len(num2) - 1; j >= 0; j -- {
			m := int(num1[i] - '0') * int(num2[j] - '0')
			sum := m + res[i + j + 1]
			res[i + j + 1] = sum % 10 
			res[i + j] +=  sum/10
		}
	}

	var sb strings.Builder 

	i := 0 

	for i < len(res) && res[i] == 0 {
		i ++
	}

	for ; i < len(res); i ++ {
		sb.WriteByte(byte(res[i] +'0'))
	}

	return sb.String()

}

