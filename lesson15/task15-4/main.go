package main

import (
	"fmt"
	"sync"
	"time"
)

var loadOnce sync.Once

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer wg.Done()
			loadOnce.Do(start)
			fmt.Println("горутина: ", i)
		}()
	}
	wg.Wait()
}

func start() {
	time.Sleep(2 * time.Second)
	fmt.Println("Slow init finished")
}
