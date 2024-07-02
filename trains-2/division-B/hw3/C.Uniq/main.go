package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	r := bufio.NewReader(os.Stdin)

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	order := make([]string, 0, 1000)
	nums := make(map[string]int, 1000)
	uniq := make([]string, 0, 1000)
	for {
		var n string
		fmt.Fscan(r, &n)
		if n == "" {
			break
		}
		if nums[n] == 0 {
			order = append(order, n)
		}

		nums[n]++

	}
	for _, n := range order {
		cnt := nums[n]
		if cnt == 1 {
			uniq = append(uniq, n)
		}
	}
	fmt.Fprintln(w, strings.Join(uniq, " "))
}
