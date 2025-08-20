package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	datachan := make(chan int)
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	defer cancel()

	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		i := 0
		defer wg.Done()
		for {
			select {
			case <-ctx.Done():
				close(datachan)
				fmt.Println("Время вышло")
				return
			case datachan <- i:
				i++
				time.Sleep(time.Millisecond * 100)

			}

		}

	}()
	wg.Add(1)
	go func() {
		defer wg.Done()
		for data := range datachan {
			fmt.Println(data)
		}
	}()
	wg.Wait()
}
