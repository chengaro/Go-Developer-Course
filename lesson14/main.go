package main

import (
	"fmt"
	"time"
)

func main() {
	result := make(chan int)
	stop := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		for i := 0; i < 3; i++ {
			result <- i
		}
		stop <- struct{}{}
	}()
	for {
		select {
		case <-stop:
			fmt.Println("завершение работы главной горутины")
			return
		case v := <-result:
			fmt.Println("результат:", v)
		default:
			fmt.Println("ожидание результата...")
			time.Sleep(time.Second)
		}
	}
}
