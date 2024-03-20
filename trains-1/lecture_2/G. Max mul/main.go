package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	list := input()
	ans1, ans2 := solution(list)
	fmt.Println(ans1, ans2)
}

func input() []int {
	list := make([]int, 0, 10)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		n := s.Text()
		i, _ := strconv.Atoi(n)
		list = append(list, i)
	}
	return list
}

func solution(list []int) (int, int) {
	if len(list) < 2 {
		return 0, 0
	}
	var max1, max2, min1, min2 int
	if list[0] > list[1] {
		max1 = 0
		max2 = 1
	} else {
		max1 = 1
		max2 = 0
	}
	min1, min2 = max2, max1

	for i := 2; i < len(list); i++ {
		if list[i] >= list[max1] {
			max2 = max1
			max1 = i
		} else {
			if list[i] >= list[max2] {
				max2 = i
			}
		}
		if list[i] <= list[min1] {
			min2 = min1
			min1 = i
		} else {
			if list[i] <= list[min2] {
				min2 = i
			}
		}
	}
	if list[min1]*list[min2] > list[max1]*list[max2] {
		return list[min1], list[min2]
	}
	return list[max2], list[max1]
}

// -6 2
// 1 5 0 -5 -2
// -1 2 -5 -3 0 6 -8 8
