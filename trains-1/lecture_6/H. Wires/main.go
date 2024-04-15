package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	k, wires := input(os.Stdin)
	ans := solution(k, wires)
	fmt.Println(ans)
}

func input(r io.Reader) (int, []int) {
	rb := bufio.NewReader(r)
	var n, k int
	fmt.Fscan(rb, &n, &k)
	wires := make([]int, n)
	for i := range wires {
		fmt.Fscan(rb, &wires[i])
	}
	return k, wires
}

func solution(k int, wires []int) int {
	l := 0
	r := 10_000_000
	slices.Sort(wires)
	check := func(l int) bool {
		return cuts(l, wires) >= k
	}
	for l < r {
		m := (l + r + 1) / 2
		if check(m) {
			l = m
		} else {
			r = m - 1
		}
	}

	return l
}

func cuts(l int, wires []int) int {
	sum := 0
	for _, w := range wires {
		sum += w / l
	}
	return sum
}
