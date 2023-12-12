package piscine

import (
	"fmt"
	"strconv"
)

/*
func main() {
	PrintNbrInOrder(321)
	PrintNbrInOrder(0)
	PrintNbrInOrder(321)
}*/

func PrintNbrInOrder(n int) {
	// string
	nstring := strconv.Itoa(n)
	nrune := []rune(nstring)

	// array to store count of each digit
	count := make([]int, 10)

	// count occurence of each digit
	for _, num := range nrune {
		count[num-'0']++
	}
	// print digits in ascending order
	for i := 0; i < 10; i++ {
		for j := 0; j < count[i]; j++ {
			fmt.Print(i)
		}
	}
}
