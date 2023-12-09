package piscine

func StrRev(s string) string {
	srune := []rune(s)                  // change string to rune
	reverse := make([]rune, len(srune)) // make a slice with length of srune to store reversed string
	for i := len(s) - 1; i >= 0; i-- {  // iterate through loop from last index item to the first
		reverse[len(srune)-1-i] = srune[i] // assigns the char at i of the original string to the corresponding in reversed slice
		// return (string(srune[i])) // print rune element at that index
	}
	return string(reverse)
}

/*
func main() {
	fmt.Printf(StrRev("fana"))
}
*/
