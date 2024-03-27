package main

import "fmt"

func main() {
	type square int
	var s square = 30
	s += square(15)
	fmt.Println(s)
}
