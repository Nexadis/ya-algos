package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

type set map[int]struct{}

func main() {
	anna, boris := input()
	res := solution(anna, boris)
	show(res)
}

func input() (set, set) {
	var n, m int
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	s.Scan()
	n, _ = strconv.Atoi(s.Text())
	s.Scan()
	m, _ = strconv.Atoi(s.Text())
	anna := input_cbs(s, n)
	boris := input_cbs(s, m)
	return anna, boris
}

func input_cbs(s *bufio.Scanner, n int) set {
	m := make(set, n)
	for i := 0; i < n; i++ {
		s.Scan()
		x, _ := strconv.Atoi(s.Text())
		m[x] = struct{}{}
	}
	return m
}

func solution(s1, s2 set) output {
	intercept := make([]int, 0, len(s1))
	left := make([]int, 0, len(s1))
	right := make([]int, 0, len(s2))
	for k := range s1 {
		if _, ok := s2[k]; ok {
			i, _ := slices.BinarySearch(intercept, k)
			intercept = slices.Insert(intercept, i, k)
		} else {
			i, _ := slices.BinarySearch(left, k)
			left = slices.Insert(left, i, k)
		}
	}
	for k := range s2 {
		if _, ok := s1[k]; !ok {
			i, _ := slices.BinarySearch(right, k)
			right = slices.Insert(right, i, k)
		}
	}
	return output{
		intercept,
		left,
		right,
	}
}

type output struct {
	intercept []int
	left      []int
	right     []int
}

func show(o output) {
	console := bufio.NewWriter(os.Stdout)
	defer console.Flush()
	fmt.Fprintln(console, len(o.intercept))
	for _, v := range o.intercept {
		fmt.Fprint(console, v, " ")
	}
	fmt.Fprintln(console)
	fmt.Fprintln(console, len(o.left))
	for _, v := range o.left {
		fmt.Fprint(console, v, " ")
	}
	fmt.Fprintln(console)
	fmt.Fprintln(console, len(o.right))
	for _, v := range o.right {
		fmt.Fprint(console, v, " ")
	}
}
