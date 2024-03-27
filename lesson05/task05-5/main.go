package main

import "fmt"

func main() {
	a := 1
	fmt.Println(a)
	change(&a)
	fmt.Println(a)
}

func change(v *int) {
	*v = *v + 1
}
