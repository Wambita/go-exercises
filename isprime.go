package main

import "fmt"

func IsPrime(nb int) bool {
	if nb <= 1 { //  for  nums less than or equal to 1 . negative nums aren't prime. 1 isnt prime
		return false
	} else if nb%2 == 0 && nb != 2 { // ensures that the only prime number divisible  by 2 is 2
		return false
	}
	for i := 2; i < nb/2; i++ { // starts loop from 2 and iterates from 2 to half the given term
		if nb%i == 0 { ////checks if n is divisible by the current loop variable value
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(IsPrime(4))
}
