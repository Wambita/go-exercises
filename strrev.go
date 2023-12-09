package piscine

func StrRev(s string) string {
	srune := []rune(s)                 // change string to rune
	for i := len(s) - 1; i >= 0; i-- { // iterate through loop from last index item to the first
		return (string(srune[i])) // print rune element at that index
	}
	return ""
}

/*
func main() {
	StrRev("fana")
}
*/
