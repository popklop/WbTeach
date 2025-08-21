package main

import "fmt"

func main() {
	mas := []string{"cat", "cat", "dog", "cat", "tree"}
	mnozh := []string{}
	mapa := make(map[string]struct{})
	for i := 0; i < len(mas); i++ {
		mapa[mas[i]] = struct{}{}
	}
	for k, _ := range mapa {
		mnozh = append(mnozh, k)
	}
	fmt.Println(mnozh)
}
