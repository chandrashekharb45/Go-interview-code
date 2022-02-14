package main

import "fmt"

func main() {
	factNum := 1
	factorial := 0
	fmt.Scanf(" Enter the number : ", &factorial)
	for i := 1; i <= 5; i++ {
		factNum = factNum * i
	}
	fmt.Printf("factorial number is %d\n", factNum)
}
