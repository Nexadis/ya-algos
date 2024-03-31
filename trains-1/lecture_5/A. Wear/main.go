package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	tshirt, pants := solution()
	fmt.Println(tshirt, pants)
}

func solution() (int, int) {
	r := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(r, &n)
	tshirts := readInts(r, n)
	fmt.Fscan(r, &m)
	pants := readInts(r, m)

	t, p := 0, 0
	sub := abs(tshirts[0] - pants[0])
	bt, bp := 0, 0

	for t < len(tshirts) && p < len(pants) {
		if abs(tshirts[t]-pants[p]) < sub {
			sub = abs(tshirts[t] - pants[p])
			bp = p
			bt = t
		}
		if tshirts[t] < pants[p] {
			t++
		} else {
			p++
		}
	}
	return tshirts[bt], pants[bp]
}

func readInts(r io.Reader, n int) []int {
	ints := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &ints[i])
	}
	return ints
}

func abs(v int) int {
	if v < 0 {
		return -v
	}
	return v
}
