package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func printStr(s string) {
	for _, r := range s {
		z01.PrintRune(r)
	}
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	/*
		pointsx := []int{(points.x)}
		pointsy := []int{(points.y)}
		printStr("x = ")
		for _, pointx := range pointsx {
			z01.PrintRune(rune(pointx))
		}
		printStr(",y = ")
		for _, pointy := range pointsy {
			z01.PrintRune(rune(pointy))
		}*/
	setPoint(points)
	printStr("x = ")
	z01.PrintRune('0' + rune(points.x/10))
	z01.PrintRune('0' + rune(points.x%10))
	printStr(",y = ")
	z01.PrintRune('0' + rune(points.y/10))
	z01.PrintRune('0' + rune(points.y%10))
	z01.PrintRune(rune(points.y))
}
