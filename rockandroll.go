package piscine

/*
func main() {
	fmt.Println(RockAndRoll(4))
	fmt.Println(RockAndRoll(9))
	fmt.Println(RockAndRoll(6))
}
*/
func RockAndRoll(n int) string {
	if n < 0 {
		return "error: number is negative \n"
	} else if n%2 == 0 && n%3 == 0 {
		return "Rock and Roll\n"
	} else if n%3 == 0 {
		return "Roll\n"
	} else if n%2 == 0 {
		return "Rock\n"
	} else {
		return "error: non divisible \n"
	}
}
