package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	rb := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(rb, &n)
	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(rb, &nums[i])
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var m int
	fmt.Fscan(rb, &m)

	for range m {
		var x int
		fmt.Fscan(rb, &x)
		l := lSearch(nums, x)
		r := rSearch(nums, x)
		if l >= r && nums[l] != x {
			l, r = -1, -1
		}
		fmt.Fprintln(w, l+1, r+1)
	}
}

func lSearch(nums []int, x int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if x > nums[m] {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func rSearch(nums []int, x int) int {
	l := 0
	r := len(nums) - 1

	for l < r {
		m := (l + r + 1) / 2
		if x < nums[m] {
			r = m - 1
		} else {
			l = m
		}
	}
	return l
}
