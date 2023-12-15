package main

import (
	"github.com/01-edu/z01"
)

type point struct { // structure
	x int
	y int
}

func printStr(s string) { // print strings
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func setPoint(ptr *point) { // assign values to the pointer structures
	ptr.x = 42
	ptr.y = 21
}

func PrintDigits(n int) { // print the numbers as runes
	if n < 0 || n > 9 {
		return
	}
	digits := []rune{48, 49, 50, 51, 52, 53, 54, 55, 56, 57} // ascii chars of 0-9
	digit := digits[n]
	z01.PrintRune(digit)
}

func PrintValue(value int) { // print numbers as runes like one by one can't print more than one num at a time
	if value < 0 {
		z01.PrintRune('-')
		value = -value
	}
	if value < 10 {
		PrintDigits(value)
	} else { // cases greater than 9
		PrintValue(value / 10)
		PrintValue(value % 10)
	}
}

func main() {
	points := &point{}
	setPoint(points)

	x, y := points.x, points.y
	printStr("x = ")
	PrintValue(x)
	printStr(", y = ")
	PrintValue(y)
	z01.PrintRune('\n')
}
