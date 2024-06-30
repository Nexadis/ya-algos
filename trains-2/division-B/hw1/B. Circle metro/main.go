package main

import "fmt"

func main() {
	var n, i, j int8

	fmt.Scan(&n, &i, &j)
	i, j = max(i, j), min(i, j)

	var ans int8
	if i-j > n-i+j {
		ans = n - i + j - 1
	} else {
		ans = i - j - 1
	}

	fmt.Println(ans)
}
