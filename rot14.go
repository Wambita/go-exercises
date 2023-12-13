package piscine

/*
func main() {
	result := Rot14("Hello! How are You?")

	for _, r := range result {
		z01.PrintRune(r)
	}
	z01.PrintRune('\n')
}*/

func Rot14(s string) string {
	a := []rune(s)
	var result string
	for _, letter := range a {
		// lowercase letters
		if letter >= 'a' && letter <= 'z' {
			result += string('a' + (letter-'a'+14)%26)
		} else if letter >= 'A' && letter <= 'Z' {
			result += string('A' + (letter-'A'+14)%26)
		} else {
			result += string(letter)
		}
	}
	return result
}
