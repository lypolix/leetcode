package main

import (
	"bufio"
	"fmt"
	"os"
)


type swap struct {i, j int}

func main() {
	in := bufio.NewReader(os.Stdin)
	out := bufio.NewWriter(os.Stdout)
	defer out.Flush()

	var n int

	if _, err := fmt.Fscan(in, &n); err != nil {
		panic(err)
	}

	tree := make([]int, n)

	for i := 0; i < n; i ++ {
		if _, err := fmt.Fscan(in, &tree[i]); err != nil {
			panic(err)
		}
	}

	swaps := buildMinHeap(tree)

	fmt.Fprintln(out, len(swaps))

	for _, s := range swaps {
		fmt.Fprintf(out, "%d %d\n", s.i, s.j)							
	}
}


func buildMinHeap(tree []int) []swap{
	n := len(tree)

	swaps := make([]swap, 0, n)

	for i := (n - 2)/2; i >= 0; i -- {
		siftDown(tree, i, n, &swaps)
	}

	return swaps 
}


func siftDown(a []int, i int, n int, swaps *[]swap) {
	for {
		l := i*2 + 1
		r := i*2 + 2
		smallest := i

		if l < n && a[l] < a[smallest] {
			smallest = l
		}

		if r < n && a[r] < a[smallest] {
			smallest = r
		}

		if smallest == i {
			return
		}

		a[i], a[smallest] = a[smallest], a[i]
		*swaps = append(*swaps, swap{i: i, j: smallest})

		i = smallest
	}
}


