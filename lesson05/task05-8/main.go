package main

import (
	"fmt"
)

type Square int

func (s Square) String() string {
	return fmt.Sprintf("%d м²", s)
}

func main() {
	var a Square = 34
	a += Square(10)

	fmt.Println(a)
}
