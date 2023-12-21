package piscine

func SplitWhiteSpace(s string) []string {
	var answer []string // store final result
	var words string    // temporarily holds char to form a word

	// iterate over each char in input string
	for i, char := range s {

		// if char is not separator append it to word variable
		if char != ' ' && char != '\t' && char != '\n' {
			words += string(char)
		}
		// if char is separator,
		if ((char == ' ' || char == '\t' || char == '\n') && words != "") || i == len(s)-1 {
			answer = append(answer, words)
			words = ""
		}
	}

	return answer
}

/*
func main() {
	s := "HelloHAhowHAareHAyou?"
	fmt.Printf("%#v\n", Split(s, "HA"))
}
*/
func Split(s, sep string) []string {
	// iteration
	for i := 0; i < len(s); i++ {
		if Indexed(s, sep) != -1 {
			s = s[0:Indexed(s, sep)] + " " + s[Indexed(s, sep)+len(sep):]
		}
	}
	return SplitWhiteSpace(s)
}

func Indexed(s string, toFind string) int {
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
