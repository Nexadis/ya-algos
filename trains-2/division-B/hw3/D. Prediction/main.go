package main

import (
	"bufio"
	"fmt"
	"os"
	"runtime/debug"
	"slices"
	"strconv"
	"strings"
)

func main() {
	debug.SetGCPercent(-1)
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	s.Scan()
	n, _ := strconv.Atoi(s.Text())

	sum := make(map[string]int, n)
	for i := 1; i <= n; i++ {
		sum[strconv.Itoa(i)] = 0
	}
	intersects := 0

	for {
		s.Scan()
		line := s.Text()
		if line == "HELP" {
			break
		}
		set := strings.Fields(line)
		s.Scan()
		line = s.Text()
		switch line {
		case "YES":
			intersects++
			sum = intersect(sum, set)
		case "NO":
			sum = sub(sum, set)
		}
	}

	ans := make([]int, 0, n)

	for k, v := range sum {
		if v == intersects {
			n, _ := strconv.Atoi(k)
			ans = append(ans, n)
		}
	}

	slices.Sort(ans)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, n := range ans {
		fmt.Fprint(w, n, " ")
	}
}

func sub(a map[string]int, b []string) map[string]int {
	for _, s := range b {
		delete(a, s)
	}
	return a
}

func intersect(a map[string]int, b []string) map[string]int {
	for _, s := range b {
		a[s]++
	}

	return a
}
