package main

import "fmt"

func main() {
	ch := make(chan string, 4)
	go func() {
		ch <- "Привет"
		ch <- "буферизованный канал!"
	}()
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
