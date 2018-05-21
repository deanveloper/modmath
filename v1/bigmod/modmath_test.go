// Each of these tests just come out of either intuition or from my notes
// in cryptography class... don't judge me
package bigmod_test


import (
	"testing"
	. "github.com/deanveloper/modmath/bigmod"
)

func TestSolve(t *testing.T) {
	NewSolveTest(5, nil, 3, 5, 10).Test(t)
	NewSolveTest(4, nil, 20, 5, 25).Test(t)
	NewSolveTest(-1, NoSolution, 20, 5, 30).Test(t)
}

func TestChineseRemainder(t *testing.T) {
	NewChineseRemainderTest(7, 2, 5, 1, 3).Test(t)
	NewChineseRemainderTest(5871, 12, 93, 29, 127).Test(t)
	NewChineseRemainderTest(1316, 5, 23, 70, 89).Test(t)
}

func TestChineseRemainderMany(t *testing.T) {
	// I only have one example in my notes of this section
	NewChineseRemainderManyTest(49, [][2]int64{{1, 3}, {4, 5}, {0, 7}}).Test(t)
}