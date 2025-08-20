package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(chanal chan int, wg *sync.WaitGroup, sigchan chan struct{}) {
	defer wg.Done()
	for {
		select {
		case chandata, ok := <-chanal:
			if !ok {
				fmt.Println("channel closed")
				return
			}
			fmt.Println(chandata)
		case <-sigchan:
			fmt.Println("Пришел сигнал отмены, остановка горутин...")
			time.Sleep(1 * time.Second)
			return
		}
	}
}

func main() {
	processFromChan(3)
}

func processFromChan(N int) {
	wg := new(sync.WaitGroup)
	sig := make(chan os.Signal, 1)
	stopWorkers := make(chan struct{})
	var i int
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)

	chanel := make(chan int)
	for i := 0; i < N; i++ {
		wg.Add(1)
		go worker(chanel, wg, stopWorkers)
	}
	for {
		select {
		case <-sig:
			fmt.Println("Остановка работающих горутин...")
			close(stopWorkers)
			close(chanel)
			wg.Wait()
			return
		default:

			chanel <- i
			i++
		}

	}

}
