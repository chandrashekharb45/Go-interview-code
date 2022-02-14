package main

import "fmt"

func main() {
	nums := []int{3, 2, 1}
	Reverse(nums)
	fmt.Println(nums)
}

func Reverse(nums []int) {
	for a, b := 0, len(nums)-1; a < b; a, b = a+1, b-1 {
		nums[a], nums[b] = nums[b], nums[a]
	}
}
