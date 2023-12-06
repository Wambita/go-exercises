package main

// fibonacci formula : Xn = Xn-1 + Xn-2
func Fibonacci(index int) int {
	if index < 0 { // catch errors  for negative numbers
		return -1
	} else if index == 0 { // return 0 as it is the first number
		return 0
	} else if index == 1 { // returns 1 as it is the second number
		return 1
	} else { // returns value at index above 1
		return Fibonacci(index-1) + Fibonacci(index-2)
	}
}

/*
func main() {
	arg1 := 4
	fmt.Println(Fibonacci(4))
}
*/
