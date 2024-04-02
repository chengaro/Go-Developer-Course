package main

import "fmt"

var cache = map[int]int64{0: 0, 1: 1}

func fib(n int) int64 {
	if f, ok := cache[n]; ok {
		return f
	} else {
		cache[n] = fib(n-2) + fib(n-1)
		return cache[n]
	}
}

func main() {
	fib(23)
	fmt.Printf("%v", cache)
}
