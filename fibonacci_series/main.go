package main

import "fmt"

func main() {
	FibonacciSeries(10)
}

func FibonacciSeries(num int) {
	x, y := 0, 1
	for i := 0; i < num; i++ {
		fmt.Printf("Series is: %d \n", x)
		x, y = y, x+y
	}
}
