package piscine

func Capitalize(s string) string {
	// capitalize functions
	ToUpper := func(s string) string {
		srune := []rune(s)
		for i, value := range s {
			if 'a' <= value && value <= 'z' {
				// convert to upper case by subtracting the ascii chars
				srune[i] = value - ('a' - 'A')
			}
		}
		result := string(srune)
		return result
	}

	ToLower := func(s string) string {
		srune := []rune(s)
		for i, value := range srune {
			if 'A' <= value && value <= 'Z' {
				// convert to lower case by adding the ascii chars
				srune[i] = value + ('a' - 'A')
			}
		}
		result := string(srune)
		return result
	}

	IsAlpha := func(s string) bool {
		// srune := []rune(s)
		for _, value := range s {
			if (value >= 'a' && value <= 'z') || (value >= 'A' && value <= 'Z') || value >= '0' && value <= '9' {
				continue
			} else {
				return false
			}
		}
		return true
	}

	s = ToLower(s)
	for i, char := range s {
		if i == 0 {
			s = ToUpper(string(char)) + s[i+1:]
		} else {
			if IsAlpha(string(char)) && !IsAlpha(string(s[i-1])) {
				if i != len(s)-1 {
					s = s[:i] + ToUpper(string(char)) + s[i+1:]
				} else {
					s = s[:1] + ToUpper(string(char))
				}
			}
		}
	}
	return s
}

/*
func main() {
	fmt.Println(Capitalize("Hello! How are you? How+are+things+4you?"))
	fmt.Println(Capitalize("yv>+L<YV^yY}J"))
	fmt.Println(Capitalize("VbE(9gjvr<_\\B"))
	fmt.Println(Capitalize("_O?H7(eusf~}y"))
}
*/
