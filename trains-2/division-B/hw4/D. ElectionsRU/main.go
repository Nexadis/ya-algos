package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type party struct {
	name   string
	id     int
	remain float64
	voices int
	got    float64
}

type ByRemain []party

func (a ByRemain) Len() int      { return len(a) }
func (a ByRemain) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a ByRemain) Less(i, j int) bool {
	if a[i].remain == a[j].remain {
		return a[j].voices < a[i].voices
	}
	return a[j].remain < a[i].remain
}

func main() {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanLines)

	allVoices := 0
	parties := make([]party, 0, 10000)

	j := 0
	for s.Scan() {
		line := s.Text()
		if line == "" {
			break
		}
		var num, name string
		for i := len(line) - 1; i >= 0; i-- {
			if line[i] == ' ' {
				name = line[:i]
				num = line[i+1:]
				break
			}
		}
		v, _ := strconv.Atoi(num)
		parties = append(parties, party{
			id:     j,
			name:   name,
			voices: v,
		})
		allVoices += v
		j++
	}

	div := float64(allVoices) / 450

	free := float64(450)

	for i, p := range parties {
		got := float64(p.voices) / div
		remain := float64(p.voices) - div*float64(int(got))

		free -= got

		parties[i].got = got
		parties[i].remain = remain
	}

	sort.Sort(ByRemain(parties))

	for free > 0 {
		for i := 0; i < len(parties) && free > 0; i++ {
			parties[i].got++
			free--
		}
	}

	sort.Slice(parties, func(i, j int) bool {
		return parties[i].id < parties[j].id
	})

	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()
	for _, p := range parties {
		fmt.Fprintln(w, p.name, int(p.got))
	}
}
