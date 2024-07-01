package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)
	s.Scan()
	s1 := strings.Fields(s.Text())
	s.Scan()
	s2 := strings.Fields(s.Text())
	set1 := make(map[string]struct{}, len(s1))
	set2 := make(map[string]struct{}, len(s2))
	for _, c := range s1 {
		set1[c] = struct{}{}
	}
	for _, c := range s2 {
		set2[c] = struct{}{}
	}

	cnt := 0
	for k := range set1 {
		if _, ok := set2[k]; ok {
			cnt++
		}
	}
	fmt.Println(cnt)
}
