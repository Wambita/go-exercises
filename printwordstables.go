package piscine

import "github.com/01-edu/z01"

func SplitWhiteSpaced(s string) []string {
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

/*
func main() {
	a := SplitWhiteSpaced("Hello how are you?")
	PrintWordsTables(a)
}
*/
func PrintWordsTables(a []string) {
	for _, word := range a {
		for _, words := range word {
			z01.PrintRune(words)
		}
		z01.PrintRune('\n')
	}
}
