package main

import (
	"log"
)

func main() {
	fruitMarket("апельсины")
	fruitMarket("груши")
	fruitMarket("персики")
}

func fruitMarket(fruit string) int {
	basket := map[string]int{
		"апельсины": 5,
		"яблоки":    3,
		"сливы":     1,
		"груши":     0,
	}
	if quantity, ok := basket[fruit]; !ok {
		log.Printf("unknown fruit: %s", fruit)
		return 0
	} else {
		return quantity
	}
}
