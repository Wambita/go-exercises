package main

import (
	"fmt"
)

func main() {
	fmt.Println(TrimAtoi("12345"))
	fmt.Println(TrimAtoi("str123ing45"))
	fmt.Println(TrimAtoi("012 345"))
	fmt.Println(TrimAtoi("Hello World!"))
	fmt.Println(TrimAtoi("sd+x1fa2W3s4"))
	fmt.Println(TrimAtoi("sd-x1fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1-fa2W3s4"))
	fmt.Println(TrimAtoi("sdx1+fa2W3s4"))
}

func TrimAtoi(s string) int {
	// trims a string and converts the remaining numeric chars to int

	// check if input string is empty, return 0
	if len(s) == 0 {
		return 0
	}

	// empty string to store numeric chars
	var num string

	// loop through each char in string range for loop
	for _, n := range s {
		// negative num: if char is - and string is empty add it to string
		if n == '-' && len(num) == 0 {
			num += "-"
		}
		// skip all chars that aren't digits
		if n < '0' && n > '9' {
			continue
		}
		// add num chars to string
		num += string(n)
	}
	// check if num is empty or only has a- , return 0 if true
	if len(num) == 0 || num == "-" {
		return 0
	}

	// convert string to int
	return Atoi(num)
}
