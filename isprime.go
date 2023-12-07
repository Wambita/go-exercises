package piscine

func IsPrime(nb int) bool {
	if nb <= 1 { //  for  nums less than or equal to 1 . negative nums aren't prime. 1 isnt prime
		return false
	}
	for i := 2; i < nb/2; i++ { //
		if nb%i == 0 { //
			return false
		}
	}
	return true
}

/*
func main() {
	fmt.Println(IsPrime(11))
}*/
