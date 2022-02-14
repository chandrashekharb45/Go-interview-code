package main

import "fmt"

func main() {
	RepeatedString("elephant", 13)
}

func RepeatedString(s string, n int64) int64 {
	inputLength := int64(len(s))
	occurrenceOfAInInput := occurrenceOfA(s)
	numberOfSubstring := int64(n / inputLength)

	numberOfAs := numberOfSubstring * occurrenceOfAInInput
	reminder := n % inputLength // a

	for i := int64(0); i < reminder; i++ {
		if s[i] == 'a' {
			numberOfAs++
		}
	}
	fmt.Println("number of a As : ", numberOfAs)
	return numberOfAs
}
func occurrenceOfA(s string) int64 {
	var count int64
	for i := range s {
		if s[i] == 'a' {
			count++
		}
	}
	fmt.Println(" occurrenceOfA ", count)
	return count
}
