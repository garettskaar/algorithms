package fibonacci

import (
	"math/big"
	"testing"
)

func TestFibDynamicNum(t *testing.T) {
	var fibs [20]*big.Int
	// First 20 fibs
	fibs[0] = big.NewInt(0)
	fibs[1] = big.NewInt(1)
	fibs[2] = big.NewInt(1)
	fibs[3] = big.NewInt(2)
	fibs[4] = big.NewInt(3)
	fibs[5] = big.NewInt(5)
	fibs[6] = big.NewInt(8)
	fibs[7] = big.NewInt(13)
	fibs[8] = big.NewInt(21)
	fibs[9] = big.NewInt(34)
	fibs[10] = big.NewInt(55)
	fibs[11] = big.NewInt(89)
	fibs[12] = big.NewInt(144)
	fibs[13] = big.NewInt(233)
	fibs[14] = big.NewInt(377)
	fibs[15] = big.NewInt(610)
	fibs[16] = big.NewInt(987)
	fibs[17] = big.NewInt(1597)
	fibs[18] = big.NewInt(2584)
	fibs[19] = big.NewInt(4181)

	for i := 0; i <= 19; i++ {
		t.Logf("Testing FibDynamicNum with n = %d", i)
		// Test 1 - 20
		got, err := FibDynamic(i)
		// Error Thrown
		if err != nil {
			t.Errorf("Error thrown! at fib(%d)", i)
		}
		// Incorrect result
		if got.Cmp(fibs[i]) != 0 {
			t.Errorf("fib(%d) = %d; want %d", i, got, fibs[i])
		}
	}
}
func TestFibDynamicNegative(t *testing.T) {
	t.Log("Testing FibDynamicNegative with n = -1")
	got, err := FibDynamic(-1)

	if got.Cmp(big.NewInt(0)) != 0 || err == nil {
		t.Errorf("FibDynamic(-1) = %d, want 0, error: %v", got, err)
	}
}

func BenchmarkFibDynamic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FibDynamic(b.N)
	}
}
