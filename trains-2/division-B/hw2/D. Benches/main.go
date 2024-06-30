package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	var l, k int
	fmt.Fscan(r, &l, &k)
	legs := make([]int, k)
	middle := (l - 1) / 2
	for i := 0; i < k; i++ {
		fmt.Fscan(r, &legs[i])
		if l%2 == 1 && legs[i] == middle {
			fmt.Println(middle)
			return
		}
	}
	left := 0
	for ; legs[left+1] <= middle; left++ {
	}
	right := k - 1
	for ; legs[right-1] > middle; right-- {
	}
	fmt.Println(legs[left], legs[right])
}
