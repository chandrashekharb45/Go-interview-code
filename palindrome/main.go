package main

import "fmt"

func main() {
	fmt.Println(isPalindrome(121))
	// fmt.Println(isStringPalindrome("abccba"))
	// fmt.Println(IsStrPalindrome("abccba"))
}

func isPalindrome(x int) bool {
	var remainder, temp, reverse int
	if x >= 0 {
		temp = x
		fmt.Println("x : ", x)
		for {
			remainder = x % 10
			reverse = reverse*10 + remainder
			x /= 10
			fmt.Println("x value : ", x)
			if x == 0 {
				break
			}
		}
		if temp == reverse {
			return true
		}
	}
	return false
}

func isStringPalindrome(str string) bool {
	reversedStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reversedStr += string(str[i])
	}
	for i := range str {
		if str[i] != reversedStr[i] {
			return false
		}
	}
	return true
}

// Time complexity : O(N)
// Space complexity : O(N)
// func IsStrPalindrome(str string) bool {
// 	result := []byte{}
// 	for i := len(str) - 1; i >= 0; i-- {
// 		result = append(result, str[i])
// 	}
// 	return str == string(result)
// }

// Time complexity : O(N)
// Space complexity : O(1)
func IsStrPalindrome(str string) bool {
	for i := 0; i < len(str); i++ {
		j := len(str) - 1 - i
		if str[i] != str[j] {
			return false
		}
	}
	return true
}
