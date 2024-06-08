package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	wg := sync.WaitGroup{}
	var stop int32 = 0

	for i := 0; i < 5; i++ {
		i := i
		wg.Add(1)
		go func() {
			defer func() {
				fmt.Println("stop горутина:", i)
				wg.Done()
			}()
			for {
				fmt.Println("сложные вычисления горутины:", i)
				time.Sleep(time.Second)
				if stop == 1 {
					return
				}
			}
		}()
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(3 * time.Second)
		fmt.Println("ой, всё!")
		atomic.AddInt32(&stop, 1)
	}()

	wg.Wait()
}
