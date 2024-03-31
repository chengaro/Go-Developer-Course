package main

import "fmt"

func main() {
	zoo := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	zoo["выдра"] = struct{}{}
	fmt.Printf("%v", zoo)
}
