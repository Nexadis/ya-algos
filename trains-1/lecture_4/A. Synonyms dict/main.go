package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() string {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	dict := make(map[string]string, n)
	for i := 0; i < n; i++ {
		s.Scan()
		word := s.Text()
		s.Scan()
		replace := s.Text()
		dict[word] = replace
		dict[replace] = word
	}
	s.Scan()
	return dict[s.Text()]
}
