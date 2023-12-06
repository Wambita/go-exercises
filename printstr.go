package piscine

import (
	"github.com/01-edu/z01"
)

func PrintStr(s string) {
	for _, str := range s {
		z01.PrintRune(str)
	}
	z01.PrintRune('\n')
}

/*
func main() {
	PrintStr("fana")
}*/
