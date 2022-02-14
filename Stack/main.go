package main

import "fmt"

// Stack :
type Stack []string

// IsEmpty :
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push :
func (s *Stack) Push(str string) {
	*s = append(*s, str)
}

// Pop :
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", false
	}

	index := len(*s) - 1   // Get index of top most element
	element := (*s)[index] // Index into slice and obtain element
	*s = (*s)[:index]      //Remove it from stack by slicing it off
	return element, true
}

func main() {
	var stack Stack
	stack.Push("this")
	stack.Push("is")
	stack.Push("Golang")
	stack.Push("Program")
	for len(stack) > 0 {
		x, y := stack.Pop()
		if y == true {
			fmt.Println(x)
		}
	}
}
