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

func solution() int {
	s := newScanner()
	n := s.ScanInt()

	blocks := make(map[int]int, n)
	for i := 0; i < n; i++ {
		w, h := s.ScanInt(), s.ScanInt()
		if blocks[w] < h {
			blocks[w] = h
		}
	}
	height := 0
	for _, h := range blocks {
		height += h
	}
	return height
}

type scanner struct {
	*bufio.Scanner
}

func newScanner() scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	return scanner{s}
}

func (s scanner) ScanInt() int {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	return n
}
