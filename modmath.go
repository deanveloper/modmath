// A package that does modular arithmetic.
//
// Doesn't use a single % operator or any math library functions.
// This is more for experience than for actual use. There is also a variation of this for big integers,
// which can be found in the bigmod package (github.com/deanveloper/modmath/bigmod)
package modmath

// Finds the least positive residue of a number
// in a given modulus
func Lpr(a, n int) int {
	c := a / n
	return a - c * n
}

// Solves the equation ax=b mod n. Note that
// if there are multiple LPR solutions that the
// lowest one is returned.
func Solve(a, b, n int) (x int) {
	gcd := gcdEuclid(a, n)

	if gcd == 0 {
		aInv, _ := eea(a, n)
		return Lpr(aInv * b, n)
	}
}