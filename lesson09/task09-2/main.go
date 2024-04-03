package main

import "fmt"

func main() {
	nums := []int{1, 2, 3}

	for _, v1 := range nums {
		fmt.Println("v1: ", v1)
		for _, v2 := range nums {
			fmt.Println("\tv2: ", v2)
			for _, v3 := range nums {
				fmt.Println("\t\tv3: ", v3)
				for _, v4 := range nums {
					fmt.Println("\t\t\tv4: ", v4)
				}
				if v3 == 1 {
					break
				}
			}
			if v2 == 1 {
				break
			}
		}
	}
}
