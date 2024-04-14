package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"slices"
)

func main() {
	ans := solution(os.Stdin)
	fmt.Println(ans)
}

func solution(in io.Reader) int {
	rb := bufio.NewReader(in)
	var n, r, c int
	fmt.Fscan(rb, &n, &r, &c)
	heights := make([]int, n)
	for i := range heights {
		fmt.Fscan(rb, &heights[i])
	}
	slices.Sort(heights)

	left := 0
	right := heights[n-1]

	check := func(diff int) bool {
		i := 0
		couples := 0
		for i <= len(heights)-c {
			// бригада не подходит
			if heights[i+c-1]-heights[i] > diff {
				i++
				// бригада с допустимой разницей
			} else {
				i += c
				couples++
			}
		}
		return couples >= r
	}

	for left < right {
		m := (left + right) / 2
		if check(m) {
			right = m
		} else {
			left = m + 1
		}
	}

	return right
}
