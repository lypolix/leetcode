package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	in := bufio.NewReader(os.Stdin)
	var n int 

	count := 0

	if _, err := fmt.Fscan(in, &n); err != nil {
		panic(err)
	}

	mat := make([][]int, n)
	for i := 0; i < n; i++ {
		row := make([]int, n)
		for j := 0; j < n; j ++ {
			if _, err := fmt.Fscan(in, &row[j]); err != nil {
				panic(err)
			}
			if row[j] == 1 {
				count += 1
			}
		}
		mat[i] = row
	}

	res := make([][]int, count)
	for i := 0; i < count; i ++ {
		res[i] = make([]int, 2)	
	}
 
	k := 0
	for i := 0; i < n; i ++ {
		for j := 0; j < n; j ++ {
			if mat[i][j] == 1{
				res[k][0] = i + 1
				res[k][1] = j + 1
				k ++
			}
		}
	}

	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	for _, num := range res {
		fmt.Fprintf(out, "%d %d\n", num[0], num[1])
	}

}