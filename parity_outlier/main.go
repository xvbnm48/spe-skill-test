package main

import "fmt"

func findoutlier(integers []int) int {
	var even, odd int
	for _, integer := range integers {
		if integer%2 == 0 {
			even++
		} else {
			odd++
		}
	}
	if even > 1 {
		return odd
	} else {
		return even
	}
}

func main() {
	integers := []int{2, 4, 0, 100, 4, 11, 2602, 36}
	if findoutlier(integers) == 11 {
		fmt.Printf("should return : %d\n", integers[6])
	} else {
		fmt.Printf("should return : %d\n", integers[3])
	}
}
