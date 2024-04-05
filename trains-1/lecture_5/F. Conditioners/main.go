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
	var n int
	fmt.Fscan(r, &n)
	conditioners := make([]int, n)
	for i := range conditioners {
		fmt.Fscan(r, &conditioners[i])
	}
	var m int
	fmt.Fscan(r, &m)
	// для каждой мощности посчитаем цену
	costByPower := make([]int, 1001)
	strongest := 0
	for i := 0; i < m; i++ {
		var p, cost int
		fmt.Fscan(r, &p, &cost)
		if costByPower[p] == 0 || cost < costByPower[p] {
			costByPower[p] = cost
		}
		if p > strongest {
			strongest = p
		}
	}
	power := strongest
	now := costByPower[power]
	for power > 0 {
		if costByPower[power] != 0 && costByPower[power] < now {
			now = costByPower[power]
		}
		costByPower[power] = now
		power--
	}
	sum := 0
	for _, c := range conditioners {
		sum += costByPower[c]
	}
	return sum
}
