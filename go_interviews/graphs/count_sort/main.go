	package main

	import (
		"bufio"
		"fmt"
		"os"
	)


	func main() {

		in := bufio.NewReader(os.Stdin)

		var n int

		if _, err := fmt.Fscan(in, &n); err != nil {
			panic(err)
		}

		mass := make([]int, n)
		
		uniq := make([]int, 10)
		for i := 0; i < n; i ++ {
			
			if _, err := fmt.Fscan(in, &mass[i]); err != nil {
				panic(err)
			}

			uniq[mass[i] - 1]++
		}

		res := make([]int, 0, n)

		

		for ii, num := range uniq {
			for i := 0; i < num; i ++ {
				res = append(res, ii + 1)
			}
		}

		out := bufio.NewWriter(os.Stdout)
		defer out.Flush()

		for i, num := range res {
			fmt.Fprintf(out, "%d", num)
			if i != n - 1 {
				fmt.Fprintf(out, " ")
			}
		}
	}