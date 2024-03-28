package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	count := solution()
	fmt.Println(count)
}

func solution() int {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	m := make(map[string]struct{}, 100)
	for s.Scan() {
		m[s.Text()] = struct{}{}
	}
	i := 0
	for range m {
		i++
	}
	return i
}
