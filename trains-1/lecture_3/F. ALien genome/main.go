package main

import "fmt"

func main() {
	g1, g2 := input()
	ans := solution(g1, g2)
	fmt.Println(ans)
}

func input() (string, string) {
	var g1, g2 string
	fmt.Scan(&g1, &g2)
	return g1, g2
}

func solution(g1, g2 string) int {
	pairs2 := make(map[string]struct{}, len(g2))
	for i := 0; i < len(g2)-1; i++ {
		pairs2[g2[i:i+2]] = struct{}{}
	}
	var ans int
	for i := 0; i < len(g1)-1; i++ {
		if _, ok := pairs2[g1[i:i+2]]; ok {
			ans += 1
		}
	}
	return ans
}
