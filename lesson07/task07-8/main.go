package main

import "fmt"

func main() {
	nums := []int{1, 2, 3, 4, 5, 6}
	nums = append(nums[:3], nums[4:]...)
	fmt.Println(nums, len(nums), cap(nums))
}
