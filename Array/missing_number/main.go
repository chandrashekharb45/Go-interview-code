package main

func main() {
	nums := []int{3, 0, 1}
	missingNumber(nums)
}

func missingNumber(nums []int) int {
	n := len(nums)
	result := 0
	sum := n * (n + 1) / 2
	for i := 0; i < n; i++ {
		result = result + nums[i]
	}
	return sum - result
}
