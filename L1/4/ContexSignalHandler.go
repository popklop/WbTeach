package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func ctxworker(ctx context.Context, datachan chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case chandata, ok := <-datachan:
			if !ok {
				fmt.Println("Channel is Closed! Returning...")
				return
			}
			fmt.Println(chandata)
		case <-ctx.Done():
			return
		}
	}

}

func main() {
	workerHandler(3)
}

func workerHandler(n int) {
	ctx := context.Background()
	wg := &sync.WaitGroup{}
	ctx, cancel := signal.NotifyContext(ctx, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	defer cancel()
	datachan := make(chan int)
	for i := 0; i < n; i++ {
		wg.Add(1)
		go ctxworker(ctx, datachan, wg)
	}

	var i int
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Stopping workers...")
			time.Sleep(500 * time.Millisecond)
			close(datachan)
			wg.Wait()
			return
		default:
			select {
			case datachan <- i:
				i++
			case <-ctx.Done():
				continue
			}
		}
	}
}
