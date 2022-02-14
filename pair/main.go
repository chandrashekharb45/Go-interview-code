package main

import "fmt"

//Pair : this function basically to count pair of elements
func main(n int32, ar []int32) int32 {
	// Write your code here
	var count int32
	m := make(map[int32]bool)

	for _, v := range ar {
		if m[v] {
			count++
			m[v] = false
		} else {
			// set true for the first item of the pair
			m[v] = true
		}
	}
	fmt.Print("pair")
	return count
}
