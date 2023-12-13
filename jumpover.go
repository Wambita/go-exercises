package piscine

/*
func main() {
	fmt.Print(JumpOver("1010101010"))
	fmt.Print(JumpOver(""))
	fmt.Print(JumpOver("t w e l v e"))
	fmt.Print(JumpOver("12"))
}*/

func JumpOver(str string) string {
	if str == "" || len(str) < 3 {
		return "\n"
	}

	result := ""
	for i := 2; i < len(str); i += 3 {
		result += string(str[i])
	}
	return result
}
