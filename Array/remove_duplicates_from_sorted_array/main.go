package main

import "fmt"

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}

func removeDuplicates(nums []int) int {
	a := 0
	if len(nums) == 0 {
		return a
	}
	for b := 1; b < len(nums); b++ {
		if nums[a] != nums[b] {
			a++
			nums[a] = nums[b]
		}
	}
	return a + 1
}
