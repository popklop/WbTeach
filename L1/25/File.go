package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func myTimer(duration time.Duration) {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, duration)
	defer cancel()
	<-ctx.Done()
	return

}

// Можно и так, но это будет доп. потреблять ресурсы ЦП
//
//	func myTimer(dursec float64) {
//		t1 := time.Now()
//		for {
//			if time.Since(t1).Seconds() > dursec {
//				return
//			}
//		}
//	}
func main() {
	wg := sync.WaitGroup{}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for {
			myTimer(200 * time.Millisecond)
			fmt.Println("doingsomething")
		}
	}()
	wg.Wait()

}
