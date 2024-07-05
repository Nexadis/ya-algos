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
	nums := make([]int, n)
	for i := range nums {
		fmt.Fscan(r, &nums[i])
	}
	slices.Sort(nums)
	fmt.Fscan(r, &n)

	ans := make([]int, 0, n)

	for range n {
		var from, to int
		fmt.Fscan(r, &from, &to)
		left := lsearch(nums, from)
		right := rsearch(nums, to)

		l := right - left + 1

		if right == left && (nums[right] > to || nums[left] < from) {
			l = 0
		}

		ans = append(ans, l)
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, a := range ans {
		fmt.Fprint(w, a, " ")
	}
}

func lsearch(nums []int, val int) (index int) {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] < val {
			l = m + 1
		} else {
			r = m
		}
	}
	return l
}

func rsearch(nums []int, val int) (index int) {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r + 1) / 2
		if nums[m] > val {
			r = m - 1
		} else {
			l = m
		}
	}
	return r
}
