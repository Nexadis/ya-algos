package main

import (
	"fmt"
)

func main() {
	var k1, m, k2, p2, n2 int
	fmt.Scan(&k1, &m, &k2, &p2, &n2)
	p1, n1 := solution(k1, m, k2, p2, n2)
	fmt.Println(p1, n1)
}

func getEntAndFloor(flatno, kpf, floors int) (int, int) {
	floorsbefore := (flatno - 1) / kpf
	entrance := floorsbefore/floors + 1
	floor := floorsbefore%floors + 1
	return entrance, floor

}

func check(k1, m, k2, p2, n2, kpf int) (int, int) {
	entrance2, floor2 := getEntAndFloor(k2, kpf, m)
	if entrance2 == p2 && floor2 == n2 {
		return getEntAndFloor(k1, kpf, m)
	}
	return -1, -1
}

func solution(k1, m, k2, p2, n2 int) (int, int) {
	ent := -1
	floor := -1
	goodflag := false
	maxval := max(k1, k2) + 1
	for kpf := 1; kpf <= maxval; kpf++ {
		nent, nfloor := check(k1, m, k2, p2, n2, kpf)
		if nent != -1 {
			goodflag = true
			if ent == -1 {
				ent, floor = nent, nfloor
			} else {
				if ent != nent && ent != 0 {
					ent = 0
				} else {
					if floor != nfloor && floor != 0 {
						floor = 0
					}
				}
			}

		}
	}
	if goodflag {
		return ent, floor
	}
	return -1, -1

}
