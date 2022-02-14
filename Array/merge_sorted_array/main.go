package main

func main() {
	num1 := []int{1, 2, 3, 0, 0, 0}
	m1 := 3
	num2 := []int{2, 5, 6}
	n1 := 3
	merge(num1, m1, num2, n1)
}

func merge(nums1 []int, m int, nums2 []int, n int) {
	if m == 0 {
		copy(nums1, nums2)
		return
	}

	i := m - 1
	j := n - 1
	k := m + n - 1

	for ; i >= 0 && j >= 0; k-- {
		if nums1[i] > nums2[j] {
			nums1[k] = nums1[i]
			i--
		} else {
			nums1[k] = nums2[j]
			j--
		}
	}

	for ; j >= 0; k-- {
		nums1[k] = nums2[j]
		j--
	}
}
