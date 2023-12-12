package piscine

/*
func main() {
	fmt.Println(ToLower("HAHAHA"))
}*/

func ToLower(s string) string {
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
