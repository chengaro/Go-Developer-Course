package main

import "fmt"

func main() {
	nums := make([]int, 0, 10)
	nums = append(nums, 4, 1, 8, 9)
	fmt.Println(nums, len(nums), cap(nums))
}
