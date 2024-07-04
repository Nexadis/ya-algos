package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rb := bufio.NewReader(os.Stdin)
	var n, q int
	fmt.Fscan(rb, &n, &q)
	sums := make([]int, n+1)
	sum := 0
	for i := 1; i <= n; i++ {
		var a int
		fmt.Fscan(rb, &a)
		sum += a
		sums[i] = sum
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i := 0; i < q; i++ {
		var l, r int
		fmt.Fscan(rb, &l, &r)
		fmt.Fprintln(w, sums[r]-sums[l-1])
	}
}
