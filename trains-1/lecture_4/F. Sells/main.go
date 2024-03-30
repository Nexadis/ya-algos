package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	buyers := solution()
	out(buyers)
}

type (
	goods  map[string]int
	buyers map[string]goods
)

func solution() buyers {
	s := bufio.NewScanner(os.Stdin)
	statistic := buyers{}
	for s.Scan() {
		statistic.scanGood(s.Text())
	}
	return statistic
}

func (b buyers) scanGood(line string) {
	s := bufio.NewScanner(strings.NewReader(line))
	s.Split(bufio.ScanWords)
	s.Scan()
	buyer := s.Text()
	s.Scan()
	good := s.Text()
	s.Scan()
	count, _ := strconv.Atoi(s.Text())
	if v := b[buyer]; v == nil {
		b[buyer] = goods{}
	}
	b[buyer][good] += count
}

func out(b buyers) {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	buyersNames := order(b)
	for _, buyer := range buyersNames {
		goods := b[buyer]
		fmt.Fprintf(w, "%s:\n", buyer)
		goodsNames := order(goods)
		for _, name := range goodsNames {
			count := goods[name]
			fmt.Fprintf(w, "%s %d\n", name, count)
		}
	}
}

func order[T any](m map[string]T) []string {
	names := make([]string, 0, len(m))
	for name := range m {
		names = append(names, name)
	}
	slices.Sort(names)
	return names
}
