package main

import (
	"bufio"
	"fmt"
	"os"
)

type tie struct{l, r int}

func main() {

	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	if _, err := fmt.Fscan(in, &n); err != nil {
		panic(err)
	}

	var m int

	if _, err := fmt.Fscan(in, &m); err != nil {
		panic(err)
	}

	ties := make([]tie, n)

	for i := 0; i < m; i ++ {
		
		var a, b int 

		if _, err := fmt.Fscan(in, &a, &b); err != nil {
			panic(err)
		}

		ties[a - 1].r ++
		ties[b - 1].l ++
	}

	for i:= 0; i < n; i ++ {
		fmt.Fprintf(out, "%d\n", ties[i].l)
		fmt.Fprintf(out, "%d\n", ties[i].r)
	}

}
