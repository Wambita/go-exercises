package piscine

func NRune(s string, n int) rune {
	srune := []rune(s)
	if n > len(s) || n <= 0 {
		return '0'
	}
	return (srune[n-1])
	return ' '
}

/*
func main() {
	NRune("Fana", 7)
}*/
