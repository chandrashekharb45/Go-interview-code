package main

func main() {
	nums := []int{1, 2, 3, 1}
	containsDuplicate(nums)
}

func containsDuplicate(nums []int) bool {
	record := make(map[int]bool, len(nums))
	for _, n := range nums {
		if _, found := record[n]; found {
			return true
		}
		record[n] = true
	}
	return false
}
