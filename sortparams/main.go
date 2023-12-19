package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]
	// sort the args
	Sort(args)
	for _, char := range args {
		for _, v := range char {
			z01.PrintRune(rune(v))
		}
		z01.PrintRune('\n')
	}
}

// sorting function

func Sort(list []string) {
	for i := 0; i < len(list)-1; i++ {
		for j := i + 1; j < len(list); j++ {
			if list[i] > list[j] {
				list[i], list[j] = list[j], list[i]
			}
		}
	}
}
