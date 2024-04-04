package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	start, end := solution(r)
	fmt.Println(start, end)
}

func solution(r io.Reader) (int, int) {
	var n, k int
	fmt.Fscan(r, &n, &k)
	trees := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var t int
		fmt.Fscan(r, &t)
		trees[i] = t
	}
	start := 1
	end := n
	kindsTrees := 0
	treesCounter := make([]int, k+1)
	treeKind := func(pos int) *int {
		return &treesCounter[trees[pos]]
	}

	left := start

	for right := start; right <= n; right++ {
		*treeKind(right)++
		if *treeKind(right) == 1 {
			kindsTrees++
		}
		for *treeKind(left) > 1 {
			*treeKind(left)--
			left++
		}
		if kindsTrees < k {
			continue
		}

		if right-left < end-start {
			start = left
			end = right
		}
	}
	return start, end
}
