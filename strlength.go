package piscine

func StrLens(s string) int {
	count := 0
	// schar := []byte(s)
	for range s {
		count = count + 1
	}
	return count
}

//use range for loop
/*
func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}*/
