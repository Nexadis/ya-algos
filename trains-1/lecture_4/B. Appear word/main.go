package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	solution(w)
}

func solution(w io.Writer) {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	occurencies := map[string]int{}
	for s.Scan() {
		fmt.Fprint(w, occurencies[s.Text()], " ")
		occurencies[s.Text()] += 1
	}
}
