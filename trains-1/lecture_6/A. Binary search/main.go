package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	answers := solution(os.Stdin)
	fmt.Println(strings.Join(answers, "\n"))
}

func solution(r io.Reader) []string {
	rd := bufio.NewReader(r)
	var n, k int
	fmt.Fscan(rd, &n, &k)
	nums := make([]int, n)
	answers := make([]string, k)
	for i := range nums {
		fmt.Fscan(rd, &nums[i])
	}

	for i := range answers {
		var find int
		fmt.Fscan(rd, &find)
		if isInclude(nums, find) {
			answers[i] = "YES"
		} else {
			answers[i] = "NO"
		}
	}
	return answers
}

func isInclude(nums []int, find int) bool {
	l := 0
	r := len(nums) - 1
	for l < r {
		m := (l + r) / 2
		if nums[m] >= find {
			r = m
		} else {
			l = m + 1
		}
	}
	return nums[l] == find
}
