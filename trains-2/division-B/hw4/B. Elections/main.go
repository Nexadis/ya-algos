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
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	order := make([]string, 0, 1000)
	candidates := make(map[string]int, 1000)
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			break
		}
		fields := strings.Fields(line)
		name, num := fields[0], fields[1]
		cnt, _ := strconv.Atoi(num)
		if _, ok := candidates[name]; !ok {
			order = append(order, name)
		}
		candidates[name] += cnt
	}
	slices.Sort(order)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	for _, name := range order {
		fmt.Fprintln(w, name, candidates[name])
	}
}
