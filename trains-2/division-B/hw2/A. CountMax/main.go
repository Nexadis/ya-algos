package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	nums := make([]int, 0, 1000)
	maxNum := -1
	for {
		var n int
		fmt.Fscan(r, &n)
		if n == 0 {
			break
		}
		if n > maxNum {
			maxNum = n
		}
		nums = append(nums, n)
	}
	cnt := 0
	for _, n := range nums {
		if n == maxNum {
			cnt++
		}
	}
	fmt.Println(cnt)
}
