package piscine

/*func main() {
	fmt.Println(Index("Hello!", "l"))
	fmt.Println(Index("Salut!", "alu"))
	fmt.Println(Index("Ola!", "hOl"))
}*/

func Index(s string, toFind string) int {
	// returns the index of the first instance of substring return -1 if not present

	// calculate lengths of input strings
	toFindLen := len(toFind)

	// iterate through each possible starting index of a substring in s
	for i := 0; i <= len(s)-len(toFind); i++ {
		// variable to track if a match is found for current index
		match := true

		// iterate through each char to find a substring
		for j := 0; j < toFindLen; j++ {
			if s[i+j] != toFind[j] { // check if theres a mismatch
				match = false
				break
			}
		}
		// if match is true return i
		if match {
			return i
		}
	}
	// return -1 if no match was found
	return -1
}
