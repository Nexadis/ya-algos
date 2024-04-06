package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	ans := solution(os.Stdin)
	fmt.Println(ans)
}

func solution(reader io.Reader) int {
	rd := bufio.NewReader(reader)
	var n, r int
	fmt.Fscan(rd, &n, &r)
	diffs := make([]int, n)
	for i := range diffs {
		fmt.Fscan(rd, &diffs[i])
	}
	sum := 0
	left := 0
	for i := 1; i < n; i++ {
		for diffs[i]-diffs[left] > r {
			left++
		}
		sum += left
	}
	return sum
}
