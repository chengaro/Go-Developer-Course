package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var cnt int
	var mu sync.Mutex
	for i := 0; i < 5; i++ {
		go func() {
			mu.Lock()
			defer mu.Unlock()
			cnt++
			fmt.Printf("Goroutine A #%d\n", cnt)
		}()
		go func() {
			mu.Lock()
			defer mu.Unlock()
			cnt++
			fmt.Printf("Goroutine B #%d\n", cnt)
		}()
	}
	time.Sleep(100 * time.Millisecond)
}
