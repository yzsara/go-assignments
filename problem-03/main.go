package main

import (
	"fmt"
	"strings"
)

func main() {

	s := "Ara S"
	finalName := compose(s, toLower, removeSpace, reverseString)
	fmt.Println(finalName)
}


func toLower(s string)string{
	return strings.ToLower(s)
}

func removeSpace(s string) string{
	return strings.ReplaceAll(s, " ", "")
}

func reverseString(s string)string{
	runes := []rune(s)
	for i, j := 0, len(runes) - 1; i < j; i, j = i+1, j-1{
		runes[i] , runes[j] = runes[j] , runes[i]
	}
	return string(runes)
}

func compose(s string, funcs ... func(string) string)string{
	for _, sampleFunc := range(funcs){
		s = sampleFunc(s)
	}
	return s
}