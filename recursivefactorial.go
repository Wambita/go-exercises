package piscine

func RecursiveFactorial(nb int) int {
	if nb == 1 || nb == 0 {
		return 1
	}
	if nb > 1 && nb < 24 {
		return nb * RecursiveFactorial(nb-1)
	} else {
		return 0
	}
}

/* recursion ==n!=n×(n−1).
//    f nb is 0 or 1, the base case is reached, and the function returns 1 (by convention, 0!=10!=1).

If nb is greater than 1, the function returns nb×RecursiveFactorial(nb−1)nb×RecursiveFactorial(nb−1), which is equivalent to nb×(nb−1)!nb×(nb−1)!.
This is a recursive call to the same function with a smaller argument (nb-1).

The recursion continues until the base case is reached, and then the results are multiplied together to compute the overall factorial.

In essence, the recursive call allows the function to break down the problem into smaller subproblems, solving each one until it reaches a base case.
This is a common approach in recursive algorithms.*/
