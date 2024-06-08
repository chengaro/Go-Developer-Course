package main

import (
	"fmt"
	"time"
)

func main() {
	go func() {
		defer fmt.Println("Прощай, дочерняя горутина уровня 1!")
		fmt.Println("Привет из дочерней горутины уровня 1!")
		go func() {
			defer fmt.Println("Прощай, дочерняя горутина уровня 2!")
			fmt.Println("Привет из дочерней горутины уровня 2!")
			go func() {
				defer fmt.Println("Прощай, дочерняя горутина уровня 3!")
				fmt.Println("Привет из дочерней горутины уровня 3!")
				go func() {
					defer fmt.Println("Прощай, дочерняя горутина уровня 4!")
					fmt.Println("Привет из дочерней горутины уровня 4!")
				}()
			}()
		}()
	}()
	time.Sleep(100 * time.Millisecond)
}
