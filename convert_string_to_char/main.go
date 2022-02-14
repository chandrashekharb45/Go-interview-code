package main

import "fmt"

func main() {
	str := "developer"
	chars := []rune(str)
	for i := 0; i < len(str); i++ {
		char := string(chars[i])
		fmt.Println(char)
	}
}
