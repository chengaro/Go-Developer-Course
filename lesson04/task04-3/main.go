package main

import "fmt"

func main() {
	greet := func() {
		fmt.Println("Hello, Go!")
	}

	hello(greet)
}

func hello(f func()) {
	f()
}
