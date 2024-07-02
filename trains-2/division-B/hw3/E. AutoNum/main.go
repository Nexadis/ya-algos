package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var m int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &m)
	witnesses := make([]map[rune]struct{}, m)
	for i := 0; i < m; i++ {
		var s string
		fmt.Fscan(r, &s)
		w := make(map[rune]struct{}, len(s))
		for _, c := range s {
			w[c] = struct{}{}
		}
		witnesses = append(witnesses, w)
	}

	var n int
	fmt.Fscan(r, &n)

	order := make([]string, 0, n)

	maxMatch := 0
	matches := make([]int, 0, n)
	for i := 0; i < n; i++ {
		var s string
		fmt.Fscan(r, &s)
		order = append(order, s)
		number := make(map[rune]struct{}, len(s))
		for _, c := range s {
			number[c] = struct{}{}
		}
		match := countMatch(number, witnesses)
		matches = append(matches, match)
		maxMatch = max(maxMatch, match)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for i, n := range order {
		if matches[i] == maxMatch {
			fmt.Fprintln(w, n)
		}
	}
}

func countMatch(number map[rune]struct{}, witnesses []map[rune]struct{}) int {
	matched := 0
nextWitness:
	for _, w := range witnesses {
		for r := range w {
			if _, ok := number[r]; !ok {
				continue nextWitness
			}
		}
		matched++
	}
	return matched
}
