package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	l, num := solution()
	fmt.Println(l, num)
}

func solution() (int, int) {
	str, k := input(os.Stdin)

	chars := make(map[rune]int, len(str))
	l := 0
	sum := 0
	pos := 0

	for r, c := range str {
		chars[c]++

		if count := chars[c]; count > k {
			for str[l] != c {
				chars[str[l]]--
				l++
			}
			chars[c]--
			l++
		}

		if sum < r-l+1 {
			sum = r - l + 1
			pos = l
		}
	}
	return sum, pos + 1
}

func input(r io.Reader) ([]rune, int) {
	rd := bufio.NewReader(r)
	var n, k int
	fmt.Fscan(rd, &n, &k)
	var s string
	fmt.Fscan(rd, &s)
	return []rune(s), k
}
