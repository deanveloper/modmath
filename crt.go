package modmath

// This file is for stuff related to the chinese remainder theorem!

// Solves x=a mod m; x=b mod n by using the chinese remainder theorem.
func SolveCrt(a, m, b, n int) int {
	s, t := eea(m, n)
	return b * s * m + a * n * t
}

// Represents an entry in the Extended Chinese Remainder Theorem
type CrtEntry struct {
	A, N int
}

// Solves the solution to x=(a1 mod m1); x=(a2 mod m2); x=...
//
// If len(eqs) == 0, it panics.
func SolveCrtMany(eqs []CrtEntry) int {
	if len(eqs) == 0 {
		return 0
	}
	if len(eqs) == 1 {
		return Lpr(eqs[0].A, eqs[0].N)
	}
	eqs2 := make([]CrtEntry, len(eqs))
	copy(eqs2, eqs)
	return solveCrtManyIntern(eqs2)
}

func solveCrtManyIntern(eqs []CrtEntry) int {
	f := eqs[0]
	s := eqs[1]
	x := SolveCrt(f.A, f.N, s.A, s.N)
	if len(eqs) == 2 {
		return x
	}
	eqs[1] = CrtEntry{x, f.N * s.N}
	return solveCrtManyIntern(eqs[1:])
}