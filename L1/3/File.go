package main

import (
	"fmt"
	"sync"
)

func worker(chanal chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for chandata := range chanal {
		fmt.Println(chandata)
	}
}

func main() {
	processFromChan(3)
}

func processFromChan(N int) {
	wg := new(sync.WaitGroup)
	chanel := make(chan int)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(chanel, wg)
	}
	for i := 0; i < 10; i++ {
		chanel <- i
	}
	close(chanel)
	wg.Wait()
	
}
