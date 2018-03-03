package modmath

// Note: this file is only for performing the euclidean algorithm
//       and extended euclidean algorithm, and should not have any
//       public-facing entries

// Finds the GCD using the euclidean algorithm
func gcdEuclid(a, b int) int {
	// first, make sure b < a
	if  b > a {
		return gcdEuclid(b, a)
	}

	// now, solve a = cb + r
	c := a / b
	r := a - c*b

	if r == 0 {
		return b
	}

	// and recurse
	return gcdEuclid(b, r)
}

// Finds x and y such that: gcdEuclid(a, b) = ax + by.
//
// The function is named because it uses the Extended Euclidean Algorithm to do this.
//
// This implementation is based on the wikibooks.org recursive python implementation.
func eea(a, b int) (x, y, gcd int) {
	if a == 0 {
		return 0, 1, b
	}
	x, y, g := eea(Lpr(b, a), a)
	return y - (b / a) * x, x, g
}