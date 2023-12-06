package piscine

func Sqrt(nb int) int {
	if nb < 0 { // catch errors  for negative numbers
		return 0
	}
	for i := 1; i <= nb; i++ {
		if i*i == nb {
			return i
		}
	}
	return 0
}

/*func main() {
	// arg1 := 4
	fmt.Println(Sqrt(4))
}*/
