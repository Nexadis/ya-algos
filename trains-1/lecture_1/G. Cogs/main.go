package main

import "fmt"

func main() {
	var n, k, m int
	fmt.Scan(&n, &k, &m)
	cogs := solution(n, k, m)
	fmt.Println(cogs)
}

func solution(n, k, m int) int {
	var cogs int
	for n >= k {
		prepares := n / k
		cog := prepares * (k / m)
		if cog == 0 {
			break
		}
		cogs += cog
		n = n%k + prepares*(k%m)
	}
	return cogs
}
