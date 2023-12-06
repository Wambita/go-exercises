package main

func UltimateDivMod(a *int, b *int) {
	var div int = *a / *b
	var mod int = *a % *b
	*a = div
	*b = mod
}

/*unc main() {
	a := 13
	b := 2
	UltimateDivMod(&a, &b)
	fmt.Println(a)
	fmt.Println(b)
}*/
