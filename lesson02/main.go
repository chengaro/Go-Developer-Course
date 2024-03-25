package main

import (
	"fmt"
	"reflect"
)

func main() {
	var r rune
	var b byte
	fmt.Printf("\nТипы rune и byte, как целочисленные типы, имеют дефолтное значение 0\nr = %d, b = %d\n\n",
		r, b)

	x := 16
	y := 3
	fmt.Printf("Результат: %d, остаток от деления: %d, тип результата: %s\n\n", x/y, x%y, reflect.TypeOf(x/y))
}
