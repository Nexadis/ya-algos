package main

import "fmt"

func main() {
	ans := solution()
	fmt.Println(ans)
}

func solution() int {
	var n int
	fmt.Scan(&n)
	seqOnes := make([]int, n+3)
	seqOnes[1] = 0
	seqOnes[2] = 0
	seqOnes[3] = 1
	for i := 4; i < n+1; i++ {
		seqOnes[i] = seqOnes[i-1] + seqOnes[i-2] + seqOnes[i-3] + 1<<(i-3)
	}
	return 1<<n - seqOnes[n]
}
