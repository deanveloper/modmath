// A package that does modular arithmetic.
package modmath

import (
	"errors"
)

var NoSolution = errors.New("no solution")

// Finds the least positive residue of a number
// in a given modulus. Note that this is very slightly
// different from the remainder (%) operator when working
// with negative numbers.
func Lpr(a, m int) int {
	return (a % m + m) % m
}

// Solves the equation `ax=b mod n` for x. Note that
// if there are multiple LPR solutions that the
// lowest one is returned. If there are no solutions,
// then (0, NoSolution) is returned
func Solve(a, b, m int) (int, error) {
	gcd := Gcd(a, m)

	// If a and m are coprime, just multiply by the inverse
	if gcd == 1 {
		aInv, _, _:= ExtendedGcd(a, m)
		aInv = Lpr(aInv, m)
		return Lpr(aInv * b, m), nil
	}

	// If gcd divides b evenly, then solve a/d x = b/d mod m/d (d = gcd)
	if Lpr(b, gcd) == 0 {
		ad := a / gcd
		bd := b / gcd
		nd := m / gcd
		return Solve(ad, bd, nd)
	}

	// else, no solution
	return 0, NoSolution
}

// Solves the equation x=a^b mod m. Note that there is not as large of a worry
// about overflowing, as a^b will not be calculated!
func SolveExp(a, b, m int) int {
	// ints contains each multiple of a^(2^i) mod m where 2^i < b
	ints := []int{a}

	for j := 2; j < b; j *= 2 {
		last := ints[len(ints) - 1]
		ints = append(ints, Lpr(last*last, m))
	}

	for i := 0; i < len(ints); i++ {
		if b &^ i == b {
			ints[i] = -1
		}
	}

	// Make a map of the powers of the ints.
	// So {7, 7, 4, 4, 7, 7} would become {7:4, 4:2}
	eq := make(map[int]int)
	for _, e := range ints {
		if e == -1 {
			continue
		}
		eq[e]++
	}

	// Simplify the map as much as possible
	modified := true
	for modified {
		modified = false
		next := make(map[int]int)

		for k, v := range eq {
			if v > 1 {
				lpr := Lpr(k*k, m)
				next[k] = eq[k] - 2
				next[lpr] = eq[lpr] + 1
				modified = true
			}
		}
		for k, v := range next {
			eq[k] = v
		}
	}

	prod := 1
	for k, v := range eq {
		for i := 0; i < v; i++ {
			prod = Lpr(prod * k, m)
		}
	}

	return prod
}