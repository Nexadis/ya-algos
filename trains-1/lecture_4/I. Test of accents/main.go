package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int {
	r := bufio.NewReader(os.Stdin)
	var n int
	fmt.Fscan(r, &n)
	dict := makeDict(r, n)
	isRightAccent := checkerAccent(dict)

	wrongs := 0

	var word string
	for {
		_, err := fmt.Fscan(r, &word)
		if err != nil {
			break
		}
		if !isRightAccent(word) {
			wrongs += 1
		}
	}
	return wrongs
}

func makeDict(r io.Reader, n int) map[string]struct{} {
	dict := make(map[string]struct{}, n)
	for i := 0; i < n; i++ {
		var word string
		fmt.Fscan(r, &word)
		dict[word] = struct{}{}
	}
	return dict
}

func checkerAccent(dict map[string]struct{}) func(w string) bool {
	words := make(map[string]struct{}, len(dict))
	for k := range dict {
		k = strings.ToLower(k)
		words[k] = struct{}{}
	}
	return func(w string) bool {
		accents := countUppers(w)
		if accents != 1 {
			return false
		}
		low := strings.ToLower(w)
		if _, ok := words[low]; !ok {
			return true
		}
		if _, ok := dict[w]; !ok {
			return false
		}
		return true
	}
}

func countUppers(w string) int {
	ups := 0
	for _, v := range w {
		if unicode.IsUpper(v) {
			ups++
		}
	}
	return ups
}
