package main

import "fmt"

const (
	_          = iota
	KB float64 = 1 << (10 * iota)
	MB
	GB
	TB
	PB
)

func main() {
	fmt.Printf("KB = %16.0f bytes\n"+
		"MB = %16.0f bytes\n"+
		"GB = %16.0f bytes\n"+
		"TB = %16.0f bytes\n"+
		"PB = %16.0f bytes", KB, MB, GB, TB, PB)
}
