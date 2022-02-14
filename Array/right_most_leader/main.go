package main

import "fmt"

// Given an array of positive integers of size N. Find the leaders in the array.
// An element of array is leader if it is greater than or equal to all the elements to its right side.
// The rightmost element is always a leader.
// Input:
// N = 6
// A[] = {16,17,4,3,5,2}
// Output: 17 5 2

func main() {
	elements := []int{16, 17, 4, 3, 5, 2}
	var j int
	for i := 0; i < len(elements); i++ {
		flag := false
		for j = i + 1; j < len(elements); j++ {
			if elements[i] <= elements[j] {
				flag = true
				break
			}
		}
		if flag == false {
			fmt.Println(elements[i])
		}
	}
}
