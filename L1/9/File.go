package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	chanint := make(chan int)
	chanx2int := make(chan int)
	mas := []int{}
	for i := 0; i < 10; i++ {
		mas = append(mas, rand.Int()%100)
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chanint)
		for c := 0; c < 10; c++ {
			chanint <- mas[c]
			time.Sleep(200 * time.Millisecond)
		}
	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(chanx2int)
		for chandata := range chanint {
			chanx2int <- chandata * 2
		}

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := range chanx2int {
			fmt.Println(i)
		}
		fmt.Println("All data read")
	}()

	wg.Wait()
}
