package bigmod

import "math/big"

// This file is for stuff related to the chinese remainder theorem!

// Solves x=a mod m; x=b mod n by using the chinese remainder theorem.
func SolveCrt(a, m, b, n *big.Int) *big.Int {
	s, t, _ := eea(m, n)

	// let eqn = bsm, eqn2 = ant
	eqn := new(big.Int)
	eqn2 := new(big.Int)
	eqn.Mul(b, s)
	eqn.Mul(eqn, m)
	eqn2.Mul(a, n)
	eqn2.Mul(eqn2, t)

	// now, let eqn = bsm + ant, eqn2 = m * n
	eqn.Add(eqn, eqn2)
	eqn2.Mul(m, n)
	return Lpr(eqn, eqn2)
}

// Represents an entry in the Extended Chinese Remainder Theorem
type CrtEntry struct {
	A, N *big.Int
}

// Solves the solution to x=(a1 mod m1); x=(a2 mod m2); x=...
//
// If len(eqs) == 0, it panics.
func SolveCrtMany(eqs []CrtEntry) *big.Int {
	if len(eqs) == 0 {
		return new(big.Int)
	}
	if len(eqs) == 1 {
		return Lpr(eqs[0].A, eqs[0].N)
	}
	eqs2 := make([]CrtEntry, len(eqs))
	copy(eqs2, eqs)
	return solveCrtManyIntern(eqs2)
}

func solveCrtManyIntern(eqs []CrtEntry) *big.Int {
	f := eqs[0]
	s := eqs[1]
	x := SolveCrt(f.A, f.N, s.A, s.N)
	if len(eqs) == 2 {
		return x
	}
	eqs[1] = CrtEntry{x, new(big.Int).Mul(f.N, s.N)}
	return solveCrtManyIntern(eqs[1:])
}