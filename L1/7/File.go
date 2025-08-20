package main

import "sync"

func main() {
	unsyncmap := make(map[int]int)
	mu := &sync.Mutex{}
	var syncmap sync.Map
	for i := 0; i < 10; i++ {
		go func() {
			k := i
			mu.Lock()
			unsyncmap[k] = 52
			mu.Unlock()
		}()
		go func() {
			k := i
			syncmap.Store(k, 52)
		}()
	}

}
