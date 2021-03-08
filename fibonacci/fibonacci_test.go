package fibonacci

import (
	"math/big"
	"testing"
)

// First 20 fibs
var fibs = [20]*big.Int{
	big.NewInt(0),
	big.NewInt(1),
	big.NewInt(1),
	big.NewInt(2),
	big.NewInt(3),
	big.NewInt(5),
	big.NewInt(8),
	big.NewInt(13),
	big.NewInt(21),
	big.NewInt(34),
	big.NewInt(55),
	big.NewInt(89),
	big.NewInt(144),
	big.NewInt(233),
	big.NewInt(377),
	big.NewInt(610),
	big.NewInt(987),
	big.NewInt(1597),
	big.NewInt(2584),
	big.NewInt(4181),
}

func TestFibDynamicNum(t *testing.T) {
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
			t.Errorf("FibDynamic(%d) = %d; want %d", i, got, fibs[i])
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
func TestFibRecursiveNum(t *testing.T) {
	var i int64
	for i = 0; i <= 19; i++ {
		t.Logf("Testing FibRecursiveNum with n = %d", i)
		// Test 1 - 20
		got := FibRecursive(i)
		// Incorrect result
		if got.Cmp(fibs[i]) != 0 {
			t.Errorf("FibRecursive(%d) = %d; want %d", i, got, fibs[i])
		}
	}
}
func TestFibRecursiveNegative(t *testing.T) {
	t.Log("Testing TestFibRecursiveNegative with n = -1")
	got := FibRecursive(-1)

	if got.Cmp(big.NewInt(0)) != 0 {
		t.Errorf("FibRecursive(-1) = %d, want 0", got)
	}
}
