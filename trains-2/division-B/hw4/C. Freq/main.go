package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type word struct {
	name string
	cnt  int
}

type ByWord []word

func (a ByWord) Len() int      { return len(a) }
func (a ByWord) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByWord) Less(i, j int) bool {
	if a[i].cnt == a[j].cnt {
		return a[j].name > a[i].name
	}
	return a[j].cnt < a[i].cnt
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	words := make(map[string]int, 1000)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			break
		}
		fields := strings.Fields(line)
		for _, word := range fields {
			words[word]++
		}
	}
	order := make([]word, 0, len(words))
	for w, cnt := range words {
		order = append(order, word{name: w, cnt: cnt})
	}
	sort.Sort(ByWord(order))
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, word := range order {
		fmt.Fprintln(w, word.name)
	}
}
