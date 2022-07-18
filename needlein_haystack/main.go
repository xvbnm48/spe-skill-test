package main

import "fmt"

func findNedle(haystack []string, needle string) int {
	for i, hay := range haystack {
		if hay == needle {
			fmt.Printf("this functionn return %d as the index of the needle\n", i)
			return i
		}
	}
	return -1
}

func main() {
	haystack := []string{"red", "blue", "yellow", "black", "grey"}
	fmt.Println(findNedle(haystack, "blue"))
}
