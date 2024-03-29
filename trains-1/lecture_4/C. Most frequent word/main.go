package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	words := input()
	ans := solution(words)
	fmt.Println(ans)
}

func input() map[string]int {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	words := map[string]int{}
	for s.Scan() {
		words[s.Text()] += 1
	}
	return words
}

func solution(words map[string]int) string {
	occurencies := 0
	val := ""
	for k, v := range words {
		if v > occurencies {
			val = k
			occurencies = v
			continue
		}
		if v == occurencies && k < val {
			val = k
		}
	}
	return val
}
