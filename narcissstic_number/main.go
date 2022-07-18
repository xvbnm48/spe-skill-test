package main

import "fmt"

func main() {
	var number, tempNumber, remainder int
	var result int = 0
	fmt.Scan(&number)

	tempNumber = number
	for {
		remainder = tempNumber % 10
		result += remainder * remainder * remainder
		tempNumber = tempNumber / 10
		if tempNumber == 0 {
			break
		}
	}

	if result == number {
		fmt.Println("True")
	} else {
		fmt.Println("False")
	}
}
