package modmath

// Note: this file is only for performing the euclidean algorithm
//       and extended euclidean algorithm

// Finds the Greatest Common Divisor using the euclidean algorithm. (Optimized to not use recursion)
func Gcd(a, b int) int {

	next := a - a/b*b
	for next != 0 {
		oldB := b
		b = next
		next = a - (a / oldB * oldB)
		a = oldB
	}
	return b
}

// Finds x and y such that: Gcd(a, b) = ax + by. (By the extended euclidean algorithm)
//
// This implementation is based on
// https://en.wikibooks.org/wiki/Algorithm_Implementation/Mathematics/Extended_Euclidean_algorithm#Iterative_algorithm_3
func ExtendedGcd(a, b int) (x, y, gcd int) {
	x0, x1, y0, y1 := 1, 0, 0, 1

	for a != 0 {
		var q int
		q, b, a = b / a, a, b % a
		x0, x1 = x1, x0 - q * x1
		y0, y1 = y1, y0 - q * y1
	}
	return y0, x0, b
}
