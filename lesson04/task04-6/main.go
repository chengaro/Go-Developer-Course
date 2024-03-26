package main

import "fmt"

/*
const (
	n    = 23
	fib0 = 0
	fib1 = 1
)

func main() {
	f := make([]int64, n)
	f[0] = fib0
	f[1] = fib1
	for i := 2; i < n; i++ {
		f[i] = f[i-2] + f[i-1]
	}
	printFib(f)
}

func printFib(f []int64) {
	for i, v := range f {
		fmt.Printf("fib%d: %d\n", i, v)
	}
}
*/

const (
	fib0  = 0
	fib1  = 1
	fib2  = fib0 + fib1
	fib3  = fib1 + fib2
	fib4  = fib2 + fib3
	fib5  = fib3 + fib4
	fib6  = fib4 + fib5
	fib7  = fib5 + fib6
	fib8  = fib6 + fib7
	fib9  = fib7 + fib8
	fib10 = fib8 + fib9
	fib11 = fib9 + fib10
	fib12 = fib10 + fib11
	fib13 = fib11 + fib12
	fib14 = fib12 + fib13
	fib15 = fib13 + fib14
	fib16 = fib14 + fib15
	fib17 = fib15 + fib16
	fib18 = fib16 + fib17
	fib19 = fib17 + fib18
	fib20 = fib18 + fib19
	fib21 = fib19 + fib20
	fib22 = fib20 + fib21
)

func main() {
	printFib(0, fib0)
	printFib(1, fib1)
	printFib(2, fib2)
	printFib(3, fib3)
	printFib(4, fib4)
	printFib(5, fib4)
	printFib(6, fib6)
	printFib(7, fib7)
	printFib(8, fib8)
	printFib(9, fib9)
	printFib(10, fib10)
	printFib(11, fib11)
	printFib(12, fib12)
	printFib(13, fib13)
	printFib(14, fib14)
	printFib(15, fib15)
	printFib(16, fib16)
	printFib(17, fib17)
	printFib(18, fib18)
	printFib(19, fib19)
	printFib(20, fib20)
	printFib(21, fib21)
	printFib(22, fib22)
}

func printFib(n int, f int64) {
	fmt.Printf("fib%d: %d\n", n, f)
}
