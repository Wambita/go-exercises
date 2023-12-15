package piscine

func Sqrt(nb int) int {
	if nb < 0 { // catch errors  for negative numbers
		return 0
	}
	for i := 0; i <= nb; i++ { // loop through number from 0 to number
		if i*i == nb { // check if square  is equal to number
			return i // return true if that is the case
		}
	}
	return 0
}

/*
func main() {
	// arg1 := 4
	fmt.Println(Sqrt(16))
}*/
