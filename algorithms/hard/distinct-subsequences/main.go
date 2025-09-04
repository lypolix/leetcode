package main

func main() {
	numDistinct("rabbbit", "rabbit")
}

func numDistinct(s string, t string) int {
	
	data := make([]int, len(t) + 1)
    data[0] = 1
	for i := 1; i <= len(s); i ++ {
		for j := len(t); j >= 1; j--{
			if s[i-1] == t[j - 1]{
				data[j] += data[j - 1]
			}
		}
	}

	return data[len(data) - 1]


}
