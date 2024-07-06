package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	rb := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(rb, &n, &k)

	points := make([]int, n)
	for i := range n {
		fmt.Fscan(rb, &points[i])
	}

	slices.Sort(points)
	l := 0
	r := points[n-1] - points[0] + 1
	a := search(l, r, func(m int) bool {
		begin := points[0] - 1
		end := begin
		lines := 0
		for _, p := range points {
			if p > end {
				begin = p
				end = p + m
				lines++
			}
		}
		return lines > k
	})
	fmt.Println(a)
}

func search(l, r int, lessMid func(m int) bool) int {
	for l < r {
		m := l + (r-l)/2
		if lessMid(m) {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}
