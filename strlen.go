package piscine

func StrLen(s string) int {
	count := 0    // initialize count var
	for range s { // loop through string
		count = count + 1 // update count
	}
	return count
}

/*
func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
	range for loop
}*/
