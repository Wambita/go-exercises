package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%#v\n", SplitWhiteSpaces("Hello how are you?"))
}

// takes a string and splits it into words based on spaces, tabs, newlines
func SplitWhiteSpaces(s string) []string {
	var answer []string // store final result
	var words string    // temporarily holds char to form a word

	// iterate over each char in input string
	for i, char := range s {

		// if char is not separator append it to word variable
		if char != ' ' && char != '\t' && char != '\n' {
			words += string(char)
		}
		// if char is separator,
		if ((char == ' ' || char == '\t' || char == '\n') && words != "") || i == len(s)-1 {
			answer = append(answer, words)
			words = ""
		}
	}

	return answer
}
