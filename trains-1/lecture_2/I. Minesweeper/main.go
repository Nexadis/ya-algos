package main

import (
	"fmt"
)

type mine struct {
	row    int
	column int
}

func main() {
	var n, m, k int
	fmt.Scan(&n, &m, &k)
	mines := inputMines(k)
	ans := solution(n, m, mines)
	out(ans)
}

func inputMines(k int) []mine {
	mines := make([]mine, k)
	for i := range mines {
		fmt.Scan(&mines[i].row, &mines[i].column)
	}
	return mines
}

func solution(n, m int, mines []mine) [][]int {
	ans := make([][]int, n)
	for i := range ans {
		row := make([]int, m)
		ans[i] = row
		for j := range ans[i] {
			ans[i][j] = getChar(i+1, j+1, mines)
		}
	}
	return ans
}

func getChar(i, j int, mines []mine) int {
	counter := 0
	for _, mine := range mines {
		if mine.row == i && mine.column == j {
			return -1
		}
		if catchRow(i, mine) && catchColumn(j, mine) {
			counter += 1
		}
	}
	return counter
}

func catchRow(i int, m mine) bool {
	switch m.row {
	case i - 1:
		fallthrough
	case i:
		fallthrough
	case i + 1:
		return true
	}
	return false
}

func catchColumn(j int, m mine) bool {
	switch m.column {
	case j - 1:
		fallthrough
	case j:
		fallthrough
	case j + 1:
		return true
	}
	return false
}

func out(ans [][]int) {
	for _, row := range ans {
		for _, elem := range row {
			if elem == -1 {
				fmt.Print("* ")
				continue
			}
			fmt.Print(elem, " ")
		}
		fmt.Println("")
	}
}
