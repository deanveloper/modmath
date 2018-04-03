// A package that does modular arithmetic.
//
// Doesn't use a single % operator or any math library functions.
// This is more for experience than for actual use. There is also a variation of this for big integers,
// which can be found in the bigmod package (github.com/deanveloper/modmath/bigmod)
package bigmod

import (
	"errors"
	"math/big"
)

var NoSolution = errors.New("no solution")

// Finds the least positive residue of a number
// in a given modulus
func Lpr(a, m *big.Int) *big.Int {
	c := new(big.Int).Div(a, m)
	r := new(big.Int).Mul(c, m)
	r.Sub(a, r)
	if r.Sign() == -1 {
		r.Add(r, m)
	}
	return r
}

// Solves the equation ax=b mod n. Note that
// if there are multiple LPR solutions that the
// lowest one is returned. If there are no solutions,
// then (0, NoSolution) is returned
func Solve(a, b, m *big.Int) (*big.Int, error) {
	gcd := Gcd(a, m)

	// If a and m are coprime, just multiply by the inverse
	if gcd.IsUint64() && gcd.Uint64() == 1 {
		aInv, _, _:= ExtendedGcd(a, m)
		aInv = Lpr(aInv, m)
		aInvB := new(big.Int).Mul(aInv, b)
		return Lpr(aInvB, m), nil
	}

	// If gcd divides b evenly, then solve a/d x = b/d mod m/d (d = gcd)
	if Lpr(b, gcd).Sign() == 0 {
		ad := new(big.Int).Div(a, gcd)
		bd := new(big.Int).Div(b, gcd)
		nd := new(big.Int).Div(m, gcd)
		return Solve(ad, bd, nd)
	}

	// else, no solution
	return new(big.Int), NoSolution
}

// Solves the equation x=a^b mod m. Note that there is not as large of a worry
// about overflowing, as a^b will not be calculated!
func SolveExp(a, b, m *big.Int) *big.Int {
	// would use log2 but I don't want to depend on math library
	ints := []*big.Int{new(big.Int).Set(a)}

	for j := big.NewInt(2); j.Cmp(b) < 0; j.Add(j, j) {
		last := ints[len(ints) - 1]
		ints = append(ints, Lpr(new(big.Int).Mul(last, last), m))
	}

	for i := 0; i < len(ints); i++ {
		if b.Bit(i) == 0 {
			ints[i] = nil
		}
	}

	// Make a map of the powers of the ints.
	// So {7, 7, 4, 4, 7, 7} would become {7:4, 4:2}
	eq := make(map[*big.Int]*big.Int)
	one := big.NewInt(1)
	for _, e := range ints {
		if e == nil {
			continue
		}
		if eq[e] == nil {
			eq[e] = new(big.Int)
		}
		eq[e].Add(eq[e], one)
	}

	// Simplify the map as much as possible
	modified := true
	for modified {
		modified = false
		next := make(map[*big.Int]*big.Int)

		for k, v := range eq {
			if v.Cmp(one) > 0 {
				next[k].Sub(next[k], one).Sub(next[k], one)
				lpr := Lpr(new(big.Int).Mul(k, k), m)
				next[lpr].Add(next[lpr], one)
				modified = true
			}
		}
		for k, v := range next {
			eq[k] = v
		}
	}

	prod := big.NewInt(1)
	for k, v := range eq {
		for i := big.NewInt(0); i.Cmp(v) < 0; i.Add(i, one) {
			prod = Lpr(new(big.Int).Mul(prod, k), m)
		}
	}

	return prod
}