package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	list := input()
	ans := solution(list)
	for _, a := range ans {
		fmt.Print(a, " ")
	}
}

func input() []int {
	list := make([]int, 0, 10)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		n, _ := strconv.Atoi(s.Text())
		list = append(list, n)
	}
	return list
}

func solution(list []int) []int {
	if len(list) < 3 {
		return nil
	}
	ansmin := make([]int, 3)
	copy(ansmin, list[:3])
	slices.Sort(ansmin)
	ansmax := make([]int, 0, 3)
	for i := len(ansmin) - 1; i >= 0; i-- {
		ansmax = append(ansmax, ansmin[i])
	}
	for i := 3; i < len(list); i++ {
		changeMax(list, ansmax, i)
		changeMin(list, ansmin, i)
	}
	ansmin[2] = ansmax[0]
	if mul(ansmax) > mul(ansmin) {
		return ansmax
	}
	return ansmin
}

func changeMax(list, maxlist []int, current int) {
	for pos, m := range maxlist {
		if list[current] >= m {
			for i := len(maxlist) - 1; i > pos; i-- {
				maxlist[i] = maxlist[i-1]
			}
			maxlist[pos] = list[current]
			break
		}
	}
}

func changeMin(list, minlist []int, current int) {
	for pos, m := range minlist {
		if list[current] <= m {
			for i := len(minlist) - 1; i > pos; i-- {
				minlist[i] = minlist[i-1]
			}
			minlist[pos] = list[current]
			break
		}
	}
}

func mul(list []int) int {
	ans := 1
	for _, v := range list {
		ans *= v
	}
	return ans
}
