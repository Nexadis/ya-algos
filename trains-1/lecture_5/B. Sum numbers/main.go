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

func solution() int {
	var n, k int
	rdr := bufio.NewReader(os.Stdin)
	fmt.Fscan(rdr, &n, &k)
	prefixes := make([]int, n+1)
	for i := 1; i < n+1; i++ {
		var num int
		fmt.Fscan(rdr, &num)
		prefixes[i] = prefixes[i-1] + num
	}
	found := 0
	l, r := 0, 0
	for l < len(prefixes) && r < len(prefixes) {
		sum := prefixes[r] - prefixes[l]
		switch {
		case sum == k:
			found++
			r++
		case sum < k:
			r++
		case sum > k:
			l++
		}

	}
	return found
}
