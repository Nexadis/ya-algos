package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var n int
	r := bufio.NewReader(os.Stdin)
	fmt.Fscan(r, &n)
	folders := make([]int, n)
	maxFolders := 0
	sumWatch := 0
	for i := 0; i < n; i++ {
		fmt.Fscan(r, &folders[i])
		sumWatch += folders[i]
		if folders[i] > maxFolders {
			maxFolders = folders[i]
		}
	}
	fmt.Println(sumWatch - maxFolders)
}
