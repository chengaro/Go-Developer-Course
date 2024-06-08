package main

import (
	"fmt"
	"sync/atomic"
	"time"
)

func main() {
	var cnt int32
	for i := 0; i < 5; i++ {
		go func() {
			atomic.AddInt32(&cnt, 1)
			fmt.Printf("goroutine A #%d\n", cnt)
		}()
		go func() {
			atomic.AddInt32(&cnt, 1)
			fmt.Printf("goroutine B #%d\n", cnt)
		}()
	}
	time.Sleep(100 * time.Millisecond)
}
