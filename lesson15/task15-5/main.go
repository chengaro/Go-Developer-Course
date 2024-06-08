package main

import (
	"fmt"
	"math/rand"
	"sync"
)

type Meteo interface {
	ReadTemp() string
	ChangeTemp(v string)
}

type Station struct {
	temp string
	mu   sync.RWMutex
}

func (s *Station) ReadTemp() string {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.temp
}

func (s *Station) ChangeTemp(v string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.temp = v
}

func main() {
	var s Station
	wg := sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		if rand.Intn(2) == 0 {
			go func() {
				defer wg.Done()
				t := rand.Intn(50)
				s.ChangeTemp(fmt.Sprintf("%d", t))
				fmt.Printf("Set temperature to %d\n", t)
			}()
		} else {
			go func() {
				defer wg.Done()
				fmt.Printf("Current temperature: %s\n", s.ReadTemp())
			}()
		}
	}
	wg.Wait()
}
