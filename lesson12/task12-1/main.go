package main

import (
	"fmt"
	"log"
)

func main() {
	a := 1
	do(a)
}

func do(v any) {
	a := 10
	if i, ok := v.(int); ok {
		a += i
	} else {
		log.Fatalln("Wrong type")
	}
	fmt.Println(a)
}
