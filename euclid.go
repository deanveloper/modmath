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
func eea(a, b int) (int, int) {
	// first, make sure b < a
	if b > a {
		return eea(b, a)
	}

	ai := a
	bi := b
	x := 0
	y := 1
	lx := 1
	ly := 0

	for bi != 0 {
		c := ai / bi
		r := bi * c - ai

		ai = b
		bi = r

		temp := x
		x = lx - c * x
		lx = temp

		temp = y
		y = ly - c * y
		ly = temp
	}

	return lx, ly
}