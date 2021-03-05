package fibonacci

import (
	"math/big"
)

func FibDynamic(n int) *big.Int {
	f := make([]*big.Int, n+1)
	f[0] = big.NewInt(0)
	f[1] = big.NewInt(1)
	for i := 2; i <= n; i++ {
		var sumPtr = big.NewInt(0)
		sumPtr.Add(f[i-1], f[i-2])
		f[i] = sumPtr
	}
	return f[n]
}
