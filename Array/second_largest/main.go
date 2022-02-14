package main

import (
	"fmt"
)

// Given an array of size N find the second largest element.
//Input:
//N = 5
//A[] = {45,17,73,89,4,79}
//Output: 79

func main() {
	elements := []int{45, 17, 73, 89, 4, 89}
	largest1 := elements[0]
	var largest2 int
	for i := 1; i < len(elements); i++ {
		if largest1 < elements[i] {
			largest2 = largest1
			largest1 = elements[i]
		} else if largest2 < elements[i] {
			largest2 = elements[i]
		}
	}
	fmt.Println("second largest element is : ", largest2)
}
