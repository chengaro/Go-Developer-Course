package main

import "fmt"

func main() {
	zoo := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}

	fmt.Printf("%+v", zoo)
}
