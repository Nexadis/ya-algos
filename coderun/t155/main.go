package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	nums := make(map[int]int, n)
	for range n {
		var x int
		fmt.Fscan(r, &x)
		nums[x]++
	}
	sum := 0
	for _, cnt := range nums {
		if cnt == 1 {
			sum++
		}
	}
	fmt.Println(sum)
}
