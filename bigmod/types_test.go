package bigmod_test

import (
	"math/big"
	"testing"
	"github.com/deanveloper/nikola"
	. "github.com/deanveloper/modmath/bigmod"
)

type SolveTest struct {
	Expected *big.Int
	ExpectedErr error
	A, B, M *big.Int
}

func NewSolveTest(exp int64, err error, a, b, m int64) SolveTest {
	var expected *big.Int = nil
	if exp != -1 {
		expected = big.NewInt(exp)
	}
	return SolveTest{expected, err, big.NewInt(a), big.NewInt(b), big.NewInt(m)}
}

func (s SolveTest) Test(t *testing.T) {
	t.Helper()
	r, e := Solve(s.A, s.B, s.M)
	if s.Expected == nil {
		nikola.SuggestTrue(t, nil == r)
	} else {
		nikola.SuggestTrue(t, s.Expected.Cmp(r) == 0)
	}
	nikola.SuggestEqual(t, s.ExpectedErr, e)
}

type SolveExpTest struct {
	Expected *big.Int
	A, B, M *big.Int
}

func NewSolveExpTest(exp int64, a, b, m int64) SolveExpTest {
	return SolveExpTest{big.NewInt(exp), big.NewInt(a), big.NewInt(b), big.NewInt(m)}
}

func (s SolveExpTest) Test(t *testing.T) {
	t.Helper()
	nikola.SuggestEqual(t, s.Expected.Text(10), SolveExp(s.A, s.B, s.M).Text(10))
}

type CrtTest struct {
	Expected *big.Int
	A, M, B, N *big.Int
}

func NewCrtTest(exp int64, a, m, b, n int64) CrtTest {
	return CrtTest{big.NewInt(exp), big.NewInt(a), big.NewInt(m), big.NewInt(b), big.NewInt(n)}
}

func (s CrtTest) Test(t *testing.T) {
	t.Helper()
	nikola.SuggestTrue(t, s.Expected.Cmp(SolveCrt(s.A, s.M, s.B, s.N)) == 0)
}

type CrtManyTest struct {
	Expected *big.Int
	Entries []CrtEntry
}

func NewCrtManyTest(exp int64, entries [][2]int64) CrtManyTest {
	arr := make([]CrtEntry, len(entries))

	for i := 0; i < len(entries); i++ {
		arr[i] = CrtEntry{A: big.NewInt(entries[i][0]), N: big.NewInt(entries[i][1])}
	}
	return CrtManyTest{big.NewInt(exp), arr}
}

func (s CrtManyTest) Test(t *testing.T) {
	t.Helper()
	nikola.SuggestTrue(t, s.Expected.Cmp(SolveCrtMany(s.Entries)) == 0)
}