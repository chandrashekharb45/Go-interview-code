package main

func main() {
	arr := []int{20, 5, 7, 25}
	search(arr, 4, 7)
}

func search(arr []int, n int, val int) int {
	for i := 0; i < n; i++ {
		if arr[i] == val {
			return i
		}
	}
	return -1
}
