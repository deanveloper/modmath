package bigmod

import "math/big"

// Note: this file is only for performing the euclidean algorithm
//       and extended euclidean algorithm, and should not have any
//       public-facing entries

// Finds the GCD using the euclidean algorithm
func Gcd(a, b *big.Int) *big.Int {
	r := new(big.Int).Div(a, b)
	r.Mul(r, b)
	r.Sub(a, r)
	if r.Sign() == 0 {
		return b
	}
	return Gcd(b, r)
}

// Finds x and y such that: Gcd(a, b) = ax + by.
//
// The function is named because it uses the Extended Euclidean Algorithm to do this.
//
// This implementation is based on the wikibooks.org recursive python implementation.
func ExtendedGcd(a, b *big.Int) (x, y, gcd *big.Int) {
	if a .Sign() == 0 {
		return new(big.Int), big.NewInt(1), b
	}
	x, y, g := ExtendedGcd(Lpr(b, a), a)

	next := new(big.Int)
	next.Div(b, a)
	next.Mul(next, x)
	next.Sub(y, next)
	return next, x, g
}