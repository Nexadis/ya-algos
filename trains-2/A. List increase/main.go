package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
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
	ans := solution(list)
	fmt.Println(ans)
}

func solution(list []int) string {
	if len(list) < 2 {
		return "YES"
	}
	prev := list[0]
	for i := 1; i < len(list); i++ {
		if list[i] <= prev {
			return "NO"
		}
		prev = list[i]
	}
	return "YES"
}
