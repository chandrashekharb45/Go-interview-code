package main

import "fmt"

func main() {
	nums := []int{3, 2, 2, 3}
	dup := removeElements(nums, 3)
	fmt.Println("No of duplicate elements : ", dup)
}

func removeElements(nums []int, val int) int {
	i, n := 0, len(nums)
	for i < n {
		if nums[i] == val {
			n--
			nums[i] = nums[n]
		} else {
			i++
		}
	}
	return n
}
