package main

import "fmt"

func main() {
	s := "some string"
	p := &s
	*p = "new string"
	fmt.Println(s)
}
