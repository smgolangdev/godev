//Implements all the exercises found in golang tour. See golang.org f//or more details.
package exers

import (
	"math"
	"strings"
)

// Calculates the squareroot of x using Newton's formula for
// finding square roots

func SquareRoot(x float64) float64 {

	switch {
	case x == 0 || math.IsNaN(x) || math.IsInf(x, 1):
		return x
	case x < 0:
		return math.NaN()
	}
	tolerance := 1e-15
	z := float64(x)

	for math.Abs(z-x/z) > (tolerance * z) {
		z = (x/z + z) / 2.0
	}
	return z
}

//////////////////////////////////////////////////////////////
//Implements the counts of each word in the given string.
//Returns a map whose key is a word and it value is the count.
//
//exercise no 43
////////////////////////////////////////////////////////////////
func WordCount(s string) map[string]int {

	arr := strings.Fields(s)
	m := make(map[string]int)

	for i := 0; i < len(arr); i++ {
		m[arr[i]] = m[arr[i]] + 1
	}
	return m
}

func fib(n int) int {
	if n < 2 {
		return 1
	}
	return fib(n-1) + fib(n-2)
}

//////////////////////////////////////////////////////////////
//Fibonacci is a function that returns a function that returns
//an int
///////////////////////////////////////////////////////////////
func Fibonacci() func() int {
	n := 0
	a, b, i := 0, 1, 0

	f := func() int {
		for i < n {
			c := a + b
			a, b = b, c
			i++
		}
		n++
		return a
	}
	return f
}

func ANew() int {
	return 20
}

///////////////////////////////////////////////////////////////////
//Modify the SqareRt function implemented earlier to return error when given a negative integer
//////////////////////////////////////////////////
type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	if e < 0 {
		return "Cannot Sqrt negative number."
	}
	return "Success"
}

func MyRt(x float64) (float64, error) {
	e := ErrNegativeSqrt(x)
	if e.Error() != "Success" {
		return x, e
	}
	return SquareRoot(x), nil
}
