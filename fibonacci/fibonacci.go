// Package fibonacci provides varies implementations of the Fibonacci sequence.
package fibonacci

import (
	"errors"
	"math/big"
)

// FibDynamic implements the fibonacci sequence with a dynamic programming approach.
// It uses big.Int to calculate very large fibonacci numbers.
func FibDynamic(n int) (*big.Int, error) {
	if n < 0 {
		return big.NewInt(0), errors.New("Negative n value")
	}

	f := make([]*big.Int, n+1)
	f[0] = big.NewInt(0)
	f[1] = big.NewInt(1)
	for i := 2; i <= n; i++ {
		var sumPtr = big.NewInt(0)
		sumPtr.Add(f[i-1], f[i-2])
		f[i] = sumPtr
	}
	return f[n], nil
}
