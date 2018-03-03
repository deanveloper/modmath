// A package that does modular arithmetic.
//
// Doesn't use a single % operator or any math library functions.
// This is more for experience than for actual use. There is also a variation of this for big integers,
// which can be found in the bigmod package (github.com/deanveloper/modmath/bigmod)
package modmath

import "errors"

var NoSolution = errors.New("no solution")

// Finds the least positive residue of a number
// in a given modulus
func Lpr(a, n int) int {
	c := a / n
	return a - c * n
}

// Solves the equation ax=b mod n. Note that
// if there are multiple LPR solutions that the
// lowest one is returned. If there are no solutions,
// then (0, NoSolution) is returned
func Solve(a, b, n int) (int, error) {
	gcd := gcdEuclid(a, n)

	// If a and n are coprime, just multiply by the inverse
	if gcd == 1 {
		aInv, _ := eea(a, n)
		return Lpr(aInv * b, n), nil
	}

	// If gcd divides b evenly, then solve a/d x = b/d mod n/d (d = gcd)
	if Lpr(b, gcd) == 0 {
		ad := a / gcd
		bd := b / gcd
		nd := n / gcd
		return Solve(ad, bd, nd)
	}

	// else, no solution
	return 0, NoSolution
}