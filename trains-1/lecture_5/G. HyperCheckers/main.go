package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int {
	rd := bufio.NewReader(os.Stdin)
	var n, k int
	fmt.Fscan(rd, &n, &k)
	uniqValues := make([]int, 0, n)
	cards := make(map[int]int, n)

	for i := 0; i < n; i++ {
		var x int
		fmt.Fscan(rd, &x)
		if _, ok := cards[x]; !ok {
			uniqValues = append(uniqValues, x)
		}
		cards[x] += 1
	}

	slices.Sort(uniqValues)
	duplicates := 0
	sum := 0
	r := 0
	for l, v := range uniqValues {

		for r < len(uniqValues) && v*k >= uniqValues[r] {

			if cards[uniqValues[r]] >= 2 {
				duplicates++
			}
			r++
		}
		rangelen := r - l
		if cards[v] >= 2 {
			sum += (rangelen - 1) * 3
		}
		if cards[v] >= 3 {
			sum++
		}
		sum += (rangelen - 1) * (rangelen - 2) * 3
		if cards[v] >= 2 {
			duplicates--
		}
		sum += duplicates * 3
	}

	return sum
}
