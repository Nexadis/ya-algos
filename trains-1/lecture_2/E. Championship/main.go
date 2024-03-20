package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	list := input(n)
	ans := solution(list)
	fmt.Println(ans)
}

func input(n int) []int {
	var a int
	list := make([]int, 0, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&a)
		list = append(list, a)
	}
	return list
}

func solution(list []int) int {
	winner := 0
	for i, v := range list {
		if v > list[winner] {
			winner = i
		}
	}
	vasiliy := 0
	for i := winner + 1; i < len(list)-1; i++ {
		if list[i]%10 != 5 {
			continue
		}
		if list[i+1] >= list[i] {
			continue
		}
		if vasiliy == 0 {
			vasiliy = i
			continue
		}
		if list[vasiliy] < list[i] {
			vasiliy = i
		}
	}
	if vasiliy == 0 {
		return 0
	}
	place := 1
	for _, v := range list {
		if v > list[vasiliy] {
			place++
		}
	}
	return place
}
