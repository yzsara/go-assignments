package main

import "fmt"

func main() {

	sara := "سارا"
	fmt.Println(sara)

	input1 := "ORDER123"
	fmt.Println(reverseString(input1))

	input2 := "CUSTOMER"
	fmt.Println(reverseString(input2))

}

func reverseString(s string) string{
	runes := []rune(s)

	for i , j := 0, len(runes) - 1; i < j; i , j = i+1, j-1{
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)

}
