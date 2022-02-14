package main

import "fmt"

func main() {
	nums := []int{2, 7, 11, 15}
	target := 9
	list := twoSum(nums, target)
	fmt.Println(list)
}

func twoSum(nums []int, target int) []int {
	hm := make(map[int]int)
	for i, left := range nums {
		// for j, right := range nums{
		//     if right + left == target && i != j{
		//         return []int{i,j}
		//     }
		// }
		_, ok := hm[left]
		if ok {
			return []int{hm[left], i}
		}
		hm[target-left] = i
	}
	return nil
}
