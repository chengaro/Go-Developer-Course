package main

import "fmt"

func main() {
	zoo := map[string]int{
		"слон":    3,
		"бегемот": 0,
		"носорог": 5,
		"лев":     1,
	}
	fmt.Println(countAnimal(zoo, "слон"))
	fmt.Println(countAnimal(zoo, "бегемот"))
	fmt.Println(countAnimal(zoo, "осьминог"))
}

func countAnimal(z map[string]int, a string) string {
	quantity, ok := z[a]
	return fmt.Sprintf("Животное: %s, количество: %d (есть в карте: %v)", a, quantity, ok)
}
