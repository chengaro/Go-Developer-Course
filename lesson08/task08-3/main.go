package main

import "fmt"

func main() {
	zoo := map[string]struct{}{
		"слон":    {},
		"бегемот": {},
		"носорог": {},
		"лев":     {},
	}
	delete(zoo, "бегемот")
	fmt.Printf("%v", zoo)
}
