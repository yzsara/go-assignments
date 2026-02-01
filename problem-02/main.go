package main

import (
	"unicode"
	"fmt"
)

func main() {

	test := "A man a plan a canal: Panama"
	fmt.Println(isPalindrome((test)))
}

func isPalindrome(s string) bool {
	runes := []rune(s)
	i, j := 0, len(runes) -1

	for i < j {
		if(!unicode.IsDigit(runes[i]) || !unicode.IsLetter(runes[i])){
			i++
			continue
		}
		if(!unicode.IsDigit(runes[j]) || !unicode.IsLetter(runes[j])){
			j--
			continue
		}
		if(unicode.ToLower(runes[i]) != unicode.ToLower(runes[j])){
			return false
		}
		i++
		j--
	}
	return true
}