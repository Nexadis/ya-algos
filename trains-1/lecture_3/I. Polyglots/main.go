package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	all, anyone := solution()
	out := bufio.NewWriter(os.Stdout)
	fmt.Fprintln(out, len(all))
	for _, l := range all {
		fmt.Fprintln(out, l)
	}
	fmt.Fprintln(out, len(anyone))
	for _, l := range anyone {
		fmt.Fprintln(out, l)
	}
	out.Flush()
}

func solution() ([]string, []string) {
	var n int
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	langs := make(map[string]int, n)
	for i := 0; i < n; i++ {
		var m int
		s.Scan()
		m, _ = strconv.Atoi(s.Text())
		for j := 0; j < m; j++ {
			s.Scan()
			langs[s.Text()] += 1
		}
	}
	all := make([]string, 0, len(langs))
	anyone := make([]string, 0, len(langs))

	for l, c := range langs {
		anyone = append(anyone, l)
		if c == n {
			all = append(all, l)
		}
	}
	return all, anyone
}
