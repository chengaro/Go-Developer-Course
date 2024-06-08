package main

import (
	"fmt"
	"sync"
)

type Config struct {
	a []int
}

// race

func _main() {
	cfg := Config{}

	go func() {
		i := 0
		for {
			i++
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			fmt.Printf("%+v\n", cfg)
			wg.Done()
		}()
	}

	wg.Wait()
}

// atomic

func main() {
	cfg := Config{}
	var lock sync.RWMutex

	go func() {
		var i int
		for {
			i++
			lock.Lock()
			cfg.a = []int{i, i + 1, i + 2, i + 3, i + 4, i + 5}
			lock.Unlock()
		}
	}()

	var wg sync.WaitGroup
	for n := 0; n < 4; n++ {
		wg.Add(1)
		go func() {
			for n := 0; n < 100; n++ {
				lock.RLock()
				fmt.Printf("%+v\n", cfg)
				lock.RUnlock()
			}
			wg.Done()
		}()
	}

	wg.Wait()
}
