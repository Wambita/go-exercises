package piscine

func StrLen(s string) int {
	count := 0
	for range s {
		count = count + 1
	}
	return count
}

/*
func main() {
	l := StrLen("Hello World!")
	fmt.Println(l)
	range for loop
}*/
