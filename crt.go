package modmath

// This file is for stuff related to the chinese remainder theorem!

// Solves x=a mod m; x=b mod n by using the chinese remainder theorem.
func SolveCrt(a, m, b, n int) int {
	s, t, _ := ExtendedGcd(m, n)
	return Lpr(b * s * m + a * n * t, m * n)
}

// Represents an entry in the Extended Chinese Remainder Theorem
type CrtEntry struct {
	A, M int
}

// Solves the solution to x=(a1 mod m1); x=(a2 mod m2); x=...
//
// If len(eqs) == 0, it panics.
func SolveCrtMany(eqs []CrtEntry) int {
	if len(eqs) == 0 {
		panic("cannot have 0 entries to solve")
	}
	if len(eqs) == 1 {
		return Lpr(eqs[0].A, eqs[0].M)
	}
	eqs2 := make([]CrtEntry, len(eqs))
	copy(eqs2, eqs)

	for i := 1; i < len(eqs2); i++ {
		x := SolveCrt(eqs2[i-1].A, eqs2[i-1].M, eqs2[i].A, eqs2[i].M)
		eqs2[i] = CrtEntry{x, eqs2[i-1].M * eqs2[i].M}
	}
	return eqs2[len(eqs2) - 1].A
}

func solveCrtManyIntern(eqs []CrtEntry) int {
	f := eqs[0]
	s := eqs[1]
	x := SolveCrt(f.A, f.M, s.A, s.M)
	if len(eqs) == 2 {
		return x
	}
	eqs[1] = CrtEntry{x, f.M * s.M}
	return solveCrtManyIntern(eqs[1:])
}