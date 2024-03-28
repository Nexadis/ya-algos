package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	nums := input()
	ans := solution(nums)
	fmt.Println(ans)
}

func input() []int {
	list := make([]int, 0, 10)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		t := s.Text()
		n, _ := strconv.Atoi(t)
		list = append(list, n)
	}
	return list
}

func solution(list []int) int {
	set := map[int]struct{}{}
	for _, v := range list {
		set[v] = struct{}{}
	}
	return len(set)
}
