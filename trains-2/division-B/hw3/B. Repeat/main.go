package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	set := make(map[string]struct{}, 1000)

	for {
		var word string
		fmt.Fscan(r, &word)
		if word == "" {
			break
		}
		if _, ok := set[word]; ok {
			fmt.Fprintln(w, "YES")
		} else {
			set[word] = struct{}{}
			fmt.Fprintln(w, "NO")
		}
	}
}
