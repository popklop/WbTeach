package main

import (
	"fmt"
	"sync"
)

type counter struct {
	val *int64
	mu  sync.Mutex
}

func main() {
	wg := sync.WaitGroup{}
	c := new(counter)
	c.val = new(int64)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		//Использование Atomic
		//go func(c *counter) {
		//	defer wg.Done()
		//	atomic.AddInt64(c.val, 1)
		//}(c)
		//Использование мьютекса
		//go func(c *counter) {
		//	defer wg.Done()
		//	c.mu.Lock()
		//	*c.val++
		//	c.mu.Unlock()
		//
		//}(c)
	}
	wg.Wait()
	fmt.Println(*c.val)
}
