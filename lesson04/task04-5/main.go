package main

import "fmt"

func main() {
	defer func() {
		fmt.Println("завершение работы")
	}()

	hello()
}

func hello() {
	fmt.Println("Hello, Go!")
}
