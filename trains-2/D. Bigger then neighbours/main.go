package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	list := input()
	ans := solution(list)
	fmt.Println(ans)
}

func input() []int {
	list := make([]int, 0, 10)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	for s.Scan() {
		i, err := strconv.Atoi(s.Text())
		if err != nil {
			continue
		}
		list = append(list, i)
	}
	return list
}

func solution(list []int) int {
	if len(list) < 3 {
		return 0
	}
	bigger := 0
	for i := 1; i < len(list)-1; i++ {
		if list[i] > list[i-1] && list[i] > list[i+1] {
			bigger += 1
		}
	}
	return bigger
}
