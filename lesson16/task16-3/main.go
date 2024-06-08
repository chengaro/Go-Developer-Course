package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	timestamp := time.Now().Add(2 * time.Second)
	ctx, cancel := context.WithDeadline(ctx, timestamp)
	defer cancel()
	wg := sync.WaitGroup{}

	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				select {
				case <-ctx.Done():
					fmt.Println("Stop goroutine", i)
					return
				default:
					fmt.Println("Hard work in goroutine", i)
					time.Sleep(time.Second)
				}
			}
		}()
	}

	wg.Wait()
}
