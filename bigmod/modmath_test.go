// Each of these tests just come out of either intuition or from my notes
// in cryptography class... don't judge me
package bigmod_test


import (
	"testing"
	. "github.com/deanveloper/modmath/bigmod"
)

func TestLpr(t *testing.T) {

	NewLprTest(0, 0, 5).Test(t)
	NewLprTest(1, 1, 10).Test(t)
	NewLprTest(2, 202, 10).Test(t)
	NewLprTest(3, 47291873, 4729187).Test(t)
}

func TestSolve(t *testing.T) {
	NewSolveTest(5, nil, 3, 5, 10).Test(t)
	NewSolveTest(4, nil, 20, 5, 25).Test(t)
	NewSolveTest(0, NoSolution, 20, 5, 30).Test(t)
}

func TestSolveExp(t *testing.T) {
	NewSolveExpTest(4, 7, 365, 9).Test(t)
	NewSolveExpTest(5, 5, 291, 11).Test(t)
}

func TestSolveCrt(t *testing.T) {
	NewCrtTest(7, 2, 5, 1, 3).Test(t)
	NewCrtTest(5871, 12, 93, 29, 127).Test(t)
	NewCrtTest(1316, 5, 23, 70, 89).Test(t)
}

func TestSolveCrtMany(t *testing.T) {
	// I only have one example in my notes of this section
	NewCrtManyTest(49, [][2]int64{{1, 3}, {4, 5}, {0, 7}}).Test(t)
}