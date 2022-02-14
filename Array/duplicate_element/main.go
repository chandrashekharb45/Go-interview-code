package main

import "fmt"

func main() {
	elements := []int{2, 5, 7, 3, 7, 4, 4}
	mp := make(map[int]int)
	for _, v := range elements {
		_, ok := mp[v]
		if ok {
			mp[v]++
			fmt.Printf("element %d is duplicate for %d times \n", v, mp[v])
		} else {
			mp[v] = 1
		}
	}
}
