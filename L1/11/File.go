package main

import "fmt"

func main() {
	mas1 := []int{1, 2, 3}
	mas2 := []int{2, 3, 4}
	finmas := []int{}
	mapa := make(map[int]struct{})
	for i := 0; i < len(mas1); i++ {
		mapa[mas1[i]] = struct{}{}
	}
	for i := 0; i < len(mas2); i++ {
		if _, exists := mapa[mas2[i]]; exists {
			finmas = append(finmas, mas2[i])
		}
	}
	fmt.Println(finmas)
}
