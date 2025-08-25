package main

import "fmt"

func main() {
	slice := make([]int, 10)
	slice = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	var index int = 3
	fmt.Scan(&index)

	copy(slice[index:], slice[index+1:])
	slice = slice[:len(slice)-1]
	fmt.Println(slice)
}
