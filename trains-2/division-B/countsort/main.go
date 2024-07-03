package main

import "fmt"

func main() {
	arr := []int{2, 2, 54, 65, 4, 4, 4, 3, 3, 4, 1, 2, 43, 54, 4, 3, 9}
	fmt.Println(arr)
	countSort(arr)
	fmt.Println(arr)
}

func countSort(arr []int) {
	maxVal := arrMax(arr)
	minVal := arrMin(arr)
	k := maxVal - minVal
	counter := make([]int, k+1)
	for _, v := range arr {
		counter[v-minVal]++
	}
	i := 0
	for j, repeat := range counter {
		for repeat > 0 {
			arr[i] = j + minVal
			i++
			repeat--
		}
	}
}

func arrMax(a []int) int {
	m := a[0]
	for _, v := range a {
		m = max(m, v)
	}
	return m
}

func arrMin(a []int) int {
	m := a[0]
	for _, v := range a {
		m = min(m, v)
	}
	return m
}
