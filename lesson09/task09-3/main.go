package main

import "fmt"

var fruit = map[string]struct{}{
	"груша":    {},
	"яблоко":   {},
	"апельсин": {},
}

var vegetable = map[string]struct{}{
	"тыква":   {},
	"огурец":  {},
	"помидор": {},
}

func main() {
	checkFood("апельсин")
	checkFood("тыква")
	checkFood("банан")
}

func checkFood(f string) {
	if _, ok := fruit[f]; ok {
		fmt.Printf("%s это фрукт\n", f)
		return
	}
	if _, ok := vegetable[f]; ok {
		fmt.Printf("%s это овощ\n", f)
		return
	}
	fmt.Printf("%s что-то странное...\n", f)
}
