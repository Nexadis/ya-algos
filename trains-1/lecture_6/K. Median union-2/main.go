package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	medians := solution(os.Stdin)
	out(os.Stdout, medians)
}

func solution(r io.Reader) []int {
	nums := input(r)
	medians := make([]int, 0, len(nums)*len(nums))
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			medians = append(medians, findMedian(nums[i], nums[j]))
		}
	}

	return medians
}

func findMedian(first, second []int) int {
	ln := len(first)
	l := min(first[0], second[0])
	r := max(first[ln-1], second[ln-1])

	for l < r {
		m := (l + r) / 2
		less := cntLess(first, m) + cntLess(second, m)
		gt := cntGt(first, m) + cntGt(second, m)
		if less <= ln-1 && gt <= ln {
			return m
		}
		if gt > ln {
			l = m + 1
		}
		if less > ln-1 {
			r = m - 1
		}
	}

	return l
}

func cntGt(nums []int, want int) int {
	return len(nums) - cntLess(nums, want+1)
}

func cntLess(nums []int, want int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] < want {
			l = m + 1
		} else {
			r = m
		}
	}
	if nums[l] < want {
		return len(nums)
	}
	return l
}

func input(r io.Reader) [][]int {
	rb := bufio.NewReader(r)
	var n, l int
	fmt.Fscan(rb, &n, &l)
	nums := make([][]int, n)
	for i := range nums {
		var x1, d1, a, c, m int
		fmt.Fscan(rb, &x1, &d1, &a, &c, &m)
		nums[i] = genSeq(x1, d1, a, c, m, l)
	}
	return nums
}

func genSeq(x, d, a, c, m, l int) []int {
	seq := make([]int, l)
	seq[0] = x
	for i := 1; i < l; i++ {
		seq[i] = x + d
		x = seq[i]
		d = ((a * d) + c) % m
	}
	return seq
}

func out(w io.Writer, medians []int) {
	wb := bufio.NewWriter(w)
	defer wb.Flush()
	for _, m := range medians {
		fmt.Fprintln(wb, m)
	}
}
