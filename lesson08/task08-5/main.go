package main

import "fmt"

func main() {
	zoo := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	zoo["бегемот"] = 2
	fmt.Printf("%v", zoo)
}
