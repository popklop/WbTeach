package main

import (
	"fmt"
	"sync"
)

func main() {
	wg := new(sync.WaitGroup)
	mas := []int{2, 4, 6, 8, 10}

	for i := 0; i < len(mas); i++ {
		wg.Add(1)
		go func(k int) {
			fmt.Println(mas[k] * mas[k])
			defer wg.Done()
		}(i)
	}
	wg.Wait()
}
