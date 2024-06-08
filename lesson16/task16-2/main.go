package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, cancel := context.WithTimeout(ctx, 2*time.Second)
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
