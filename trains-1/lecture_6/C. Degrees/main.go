package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int64 {
	rd := bufio.NewReader(os.Stdin)
	var w, h, n int64
	fmt.Fscan(rd, &w, &h, &n)
	var l int64 = 0
	r := max(w, h) * n
	for l < r {
		m := (l + r) / 2
		caph := m / h
		capw := m / w
		if caph*capw >= n {
			r = m
		} else {
			l = m + 1
		}
	}
	return l
}
