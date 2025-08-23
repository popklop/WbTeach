package main

import (
	"fmt"
)

func main() {
	sortedarr := []int{1, 3, 5, 7, 9, 11, 13, 15}
	fmt.Println(binsearch(sortedarr, 5))
}

func binsearch(arr []int, target int) int {
	fmt.Println(arr)
	l := 0
	r := len(arr) - 1
	for l <= r {
		mid := (l + r) / 2
		res := arr[mid]
		if target == res {
			return mid
		}
		if target < res {
			r = mid - 1
		} else {
			l = mid + 1
		}
	}
	return -1
}
