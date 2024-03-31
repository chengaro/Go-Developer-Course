package main

import "fmt"

func main() {
	nums1 := []int{1, 2, 3}
	nums2 := []int{4, 5}
	nums := append(nums1, nums2...)
	fmt.Println(nums, len(nums), cap(nums))
}
