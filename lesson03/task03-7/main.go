package main

import "fmt"

const (
	un float64 = 1 << iota
	deux
	quatre
	huit
	seize
)

func main() {
	fmt.Println(un, deux, quatre, huit, seize)
}
