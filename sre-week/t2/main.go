package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type packet struct {
	name string
	deps []string
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	packets := parseDigraph(s)
	installed := make(map[string]struct{}, len(packets))
	order := make([]string, 0, len(packets))

	for len(order) != len(packets) {
		for _, p := range packets {
			if _, ok := installed[p.name]; ok {
				continue
			}
			if len(p.deps) == 0 {
				installed[p.name] = struct{}{}
				order = append(order, p.name)
				continue
			}
			depCnt := 0
			for _, d := range p.deps {
				if _, ok := installed[d]; ok {
					depCnt++
				}
			}

			if depCnt == len(p.deps) {
				installed[p.name] = struct{}{}
				order = append(order, p.name)
			}
		}
	}

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, p := range order {
		fmt.Fprintln(w, p)
	}
}

func parseDigraph(s *bufio.Scanner) map[string]packet {
	s.Split(bufio.ScanLines)
	packets := make(map[string]packet, 10)
	for s.Scan() {
		line := s.Text()
		if !strings.Contains(line, "->") {
			continue
		}
		line = strings.TrimSpace(line)
		words := strings.SplitN(line, " ", 4)
		if len(words) < 3 {
			continue
		}
		name := strings.Trim(words[0], "\"")
		dep := strings.Trim(words[2], "\"")
		if _, ok := packets[name]; !ok {
			packets[name] = packet{name: name, deps: make([]string, 0, 5)}
		}
		p := packets[name]
		p.deps = append(p.deps, dep)
		packets[name] = p
		if _, ok := packets[dep]; !ok {
			packets[dep] = packet{name: dep}
		}
	}

	return packets
}
