package main

import (
	"fmt"
)

func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}

func Index(s string, toFind string) int {
	// returns the index of the first instance of substring return -1 if not present
	srune := []rune(s)
	toFindRune := []rune(toFind)

	for i, value := range srune {
		if value == toFindRune[0] {
			return i
		}
		if value != toFindRune[0] {
		}
	}
	return -1
}
