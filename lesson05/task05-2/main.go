package main

import "fmt"

func main() {
	s := "some string"
	fmt.Printf("value: %s, pointer: %v", s, &s)
}
