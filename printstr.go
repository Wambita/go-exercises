package piscine

import "fmt"

func PrintStr(s string) {
	for _, str := range s {
		fmt.Printf("%c", str)
	}
	fmt.Println("")
}

/*
func main() {
	PrintStr("fana")
}*/
