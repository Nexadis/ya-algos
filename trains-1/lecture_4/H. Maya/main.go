package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int {
	r := bufio.NewReader(os.Stdin)
	var lword, lseq int
	fmt.Fscan(r, &lword, &lseq)
	var word, seq string
	fmt.Fscan(r, &word, &seq)
	wDict := makeDict(word)
	sDict := makeDict(seq[0:lword])

	matchingLetters := matchDicts(wDict, sDict)

	same := 0
	if matchingLetters == len(wDict) {
		same += 1
	}

	for i := lword; i < lseq; i++ {
		matchingLetters += modifyDict(sDict, wDict, seq[i-lword], -1)
		matchingLetters += modifyDict(sDict, wDict, seq[i], 1)
		if matchingLetters == len(wDict) {
			same += 1
		}
	}
	return same
}

func makeDict(s string) map[byte]int {
	m := map[byte]int{}
	for _, v := range s {
		m[byte(v)] += 1
	}
	return m
}

func matchDicts(m1, m2 map[byte]int) int {
	matched := 0
	for k, v := range m1 {
		if v2 := m2[k]; v == v2 {
			matched += 1
		}
	}
	return matched
}

func modifyDict(sDict, wDict map[byte]int, symb byte, changeMatching int) int {
	matched := 0
	if _, ok := wDict[symb]; ok && wDict[symb] == sDict[symb] {
		matched = -1
	}
	sDict[symb] += changeMatching
	if _, ok := wDict[symb]; ok && wDict[symb] == sDict[symb] {
		matched = 1
	}
	return matched
}
