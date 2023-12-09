package piscine

/*
import (
	"fmt"
)

func main() {
	a := 0
	b := 1
	Swap(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}
*/
func Swap(a *int, b *int) { // func takes in two  pointer variables to swap their values
	swap := *a // introduce a temporary variable to store value pointed to by a
	*a = *b    // store the value pointed to by b in pointer a
	*b = swap  // stores the value in swap in pointer b
}
