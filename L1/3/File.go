package main

import "fmt"

func worker(chanal chan int) {
	for chandata := range chanal {
		fmt.Println(chandata)
	}
}

func main() {
	processFromChan(3)
}

func processFromChan(N int) {
	chanel := make(chan int)
	for i := 0; i < 3; i++ {
		go worker(chanel)
	}
	for i := 0; i < 10; i++ {
		chanel <- i
	}
	close(chanel)
}
