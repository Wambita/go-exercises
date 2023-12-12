package piscine

/*
func main() {
	fmt.Println(ToUpper("Hello! How are you?"))
}
*/
func ToUpper(s string) string {
	srune := []rune(s)
	for i, value := range srune {
		if 'a' <= value && value <= 'z' {
			// convert to upper case by subtracting the ascii chars
			srune[i] = value - ('a' - 'A')
		}
	}
	result := string(srune)
	return result
}
