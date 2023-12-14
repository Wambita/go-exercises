package piscine

/*
func main() {
	fmt.Println(ActiveBits(4))
}
*/
// func  counts the no. of active bits // value ==1
func ActiveBits(n int) int {
	var countactive int // initialize var to keep track of the count of active bits
	for n > 0 {         // loop until n becomes 0
		if n%2 == 1 { // check if the rightmost bit is 1 (n%2 == 1)
			countactive++ // increment count if rightmost bit is 1
		}
		n = n / 2 // shifts bits to the right by dividing n by 2 (repeat process till n becomes 0 )
	}
	return countactive // return the final count of active bits
}
