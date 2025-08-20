package main

import (
	"context"
	"fmt"
	"runtime"
	"sync"
	"time"
)

func main() {
	someCondition := true
	wg := &sync.WaitGroup{}
	sigchan := make(chan struct{})
	ctx := context.Background()

	ExpressiСonExit(&someCondition, wg)
	time.Sleep(1 * time.Second)
	someCondition = false
	time.Sleep(1 * time.Second)

	ChanNitfy(sigchan, wg)
	time.Sleep(3 * time.Second)
	fmt.Println("Time is come")
	sigchan <- struct{}{}
	close(sigchan)

	ctx, cancel := context.WithTimeout(ctx, time.Second*3)
	byContext(ctx, wg)
	defer cancel()
	time.Sleep(5 * time.Second)

	byGoexit(wg)

	wg.Wait()

}

func ExpressiСonExit(someCondition *bool, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for *someCondition {
			fmt.Println("Doing Something...")
			time.Sleep(200 * time.Millisecond)
		}
		fmt.Println("Stopping the Work...")
		defer wg.Done()
	}()

}

func ChanNitfy(stop chan struct{}, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-stop:
				fmt.Println("Exiting...")
				wg.Done()
				return
			default:
				fmt.Println("Doing something...")
				time.Sleep(200 * time.Millisecond)
			}
		}

	}()

}

func byContext(ctx context.Context, wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		for {
			select {
			case <-ctx.Done():
				fmt.Println("Exiting by context...")
				wg.Done()
				return

			default:
				fmt.Println("Doing something3...")
				time.Sleep(200 * time.Millisecond)
			}
		}

	}()
}

func byGoexit(wg *sync.WaitGroup) {
	wg.Add(1)
	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			time.Sleep(200 * time.Millisecond)
			fmt.Println("Doing something4...")
			if i == 9 {
				fmt.Println("Немедленно завершаем горутину!")
				runtime.Goexit()
			}

		}
	}()
}
