package piscine

/*
func main() {
	x := 5
	y := &x
	z := &y
	a := &z

	w := 2
	b := &w

	u := 7
	e := &u
	f := &e
	g := &f
	h := &g
	i := &h
	j := &i
	c := &j

	k := 6
	l := &k
	m := &l
	n := &m
	d := &n

	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)

	Enigma(a, b, c, d)

	fmt.Println("After using Enigma")
	fmt.Println(***a)
	fmt.Println(*b)
	fmt.Println(*******c)
	fmt.Println(****d)
}
*/
func Enigma(a ***int, b *int, c *******int, d ****int) {
	e := ***a        // take content of a store in e a is empty
	***a = *b        // b contents store in a b is empty
	*b = ****d       // d contents store in  b , d is empty
	****d = *******c // c contents store in d , c is empty
	*******c = e     // e contents store in c e is empty
}
