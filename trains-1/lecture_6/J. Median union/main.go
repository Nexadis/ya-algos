package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	leftMedians := solution(os.Stdin)
	out(os.Stdout, leftMedians)
}

func solution(r io.Reader) []int {
	rb := bufio.NewReader(r)
	var n, l int
	fmt.Fscan(rb, &n, &l)
	medians := make([][]int, n)
	for i := range medians {
		medians[i] = make([]int, l)
		for j := range medians[i] {
			fmt.Fscan(rb, &medians[i][j])
		}
	}

	leftMedians := make([]int, 0, n*n)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			leftMedians = append(leftMedians, findMedian(medians[i], medians[j]))
		}
	}
	return leftMedians
}

func findMedian(first, second []int) int {
	l1 := 0
	l2 := 0
	cnt := 0
	for l1 < len(first) && l2 < len(second) {
		cnt++
		if first[l1] < second[l2] {
			if cnt == len(first) {
				return first[l1]
			}
			l1++
		} else {
			if cnt == len(second) {
				return second[l2]
			}
			l2++
		}
	}

	return 0
}

func out(w io.Writer, nums []int) {
	wb := bufio.NewWriter(w)
	defer wb.Flush()
	for _, n := range nums {
		fmt.Fprintln(wb, n)
	}
}
