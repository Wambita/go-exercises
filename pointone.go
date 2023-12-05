package piscine

/*import (
	"fmt"
)

func main() {
	n := 0
	PointOne(&n)   // calls the address of n
	fmt.Println(n) // prints the value on the address
}*/

func PointOne(n *int) {
	*n = 1 // stores the address of 1
}
