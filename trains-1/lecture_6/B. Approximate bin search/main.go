package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ans := solution()
	out(ans)
}

func solution() []int {
	r := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(r, &n, &k)
	seq := make([]int, n)
	search := make([]int, k)
	for i := range seq {
		fmt.Fscan(r, &seq[i])
	}
	for i := range search {
		var find int
		fmt.Fscan(r, &find)
		search[i] = nearest(seq, find)
	}
	return search
}

func nearest(seq []int, find int) int {
	l := 0
	r := len(seq) - 1
	for l+1 < r {
		m := (r + l) / 2
		if seq[m] >= find {
			r = m
		} else {
			l = m
		}
	}
	if find-seq[l] <= seq[r]-find {
		return seq[l]
	}
	return seq[r]
}

func out(nums []int) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, n := range nums {
		fmt.Fprintln(w, n)
	}
}
