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
	r := bufio.NewReader(os.Stdin)
	var memSize int
	fmt.Fscan(r, &memSize)
	var compile string
	fmt.Fscan(r, &compile)
	streak := 0
	sum := 0
	for actPos := memSize; actPos < len(compile); actPos++ {
		if compile[actPos] == compile[actPos-memSize] {
			streak++
		} else {
			streak = 0
		}
		sum += streak
	}
	return sum
}
