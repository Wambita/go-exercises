package main

func StrLen(s string) int {
	count := 0
	// schar := []byte(s)
	for range s {
		count = count + 1
	}
	return count
}

/*
func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
}*/
