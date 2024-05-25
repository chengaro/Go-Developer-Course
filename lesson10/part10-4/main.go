package main

import "fmt"

func main() {
	f := fib()
	for x, n := f(); n <= 23; x, n = f() {
		fmt.Printf("fib%d: %d\n", n, x)
	}
}

func fib() func() (int, int) {
	a, b, n := 0, 1, 0
	return func() (int, int) {
		a, b = b, a+b
		n++
		return a, n
	}
}
