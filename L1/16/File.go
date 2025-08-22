package main

import "fmt"

func main() {
	fmt.Println(qsort([]int{4, 1, 2, 6, 1, 5}))
}

func qsort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	mid := len(arr) / 2
	larr := []int{}
	rarr := []int{}
	for i := 0; i < len(arr); i++ {
		if arr[i] == mid {
			continue
		}
		if arr[i] < arr[mid] {
			larr = append(larr, arr[i])
		} else {
			rarr = append(rarr, arr[i])
		}
	}
	sortedLeft := qsort(larr)
	sortedRight := qsort(rarr)
	return append(append(sortedLeft, arr[mid]), sortedRight...)
}
