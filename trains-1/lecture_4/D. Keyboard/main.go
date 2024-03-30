package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	solution(w)
}

func solution(w io.Writer) {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	keys := make([]int, n+1)
	for i := 1; i <= n; i++ {
		s.Scan()
		limit, _ := strconv.Atoi(s.Text())
		keys[i] = limit
	}
	s.Scan()
	k, _ := strconv.Atoi(s.Text())

	for i := 0; i < k; i++ {
		s.Scan()
		pushed, _ := strconv.Atoi(s.Text())
		keys[pushed] -= 1

	}
	for i := 1; i <= n; i++ {
		if keys[i] < 0 {
			fmt.Fprintln(w, "YES")
			continue
		}
		fmt.Fprintln(w, "NO")
	}
}
