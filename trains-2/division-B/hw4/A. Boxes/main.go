package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	colors := make(map[int64]int64, n)
	order := make([]int64, 0, n)

	for i := 0; i < n; i++ {
		var d, a int64
		fmt.Fscan(r, &d, &a)
		if _, ok := colors[d]; !ok {
			order = append(order, d)
		}
		colors[d] += a
	}
	slices.Sort(order)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, d := range order {
		fmt.Fprintln(w, d, colors[d])
	}
}
