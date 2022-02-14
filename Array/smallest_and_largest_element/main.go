package main

import "fmt"

func main() {
	elements := []int{12, 56, 76, 89, 100, 343, 21, 234}
	largest := elements[0]
	smallest := elements[0]
	for i := 0; i < len(elements); i++ {
		if elements[i] > largest {
			largest = elements[i]
		} else if elements[i] < smallest {
			smallest = elements[i]
		}
	}

	fmt.Printf("smallest element is %d\n", smallest)

	fmt.Printf("Largest element is %d\n", largest)
}
