package main

import "fmt"

func main() {
	fruit := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	fruit[2] = "персик"
	fmt.Println(fruit)
}
