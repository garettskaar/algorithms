package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"time"
)

func main() {

	nPtr := flag.Int("n", 42, "fibonacci number to calculate")

	flag.Parse()

	if *nPtr > 999999 {
		for {
			var response string
			fmt.Println("Warning, it may take several minutes to calculate this fibonacci number.\n\nContinue? Enter Y or N")
			fmt.Scan(&response)
			if response == "Y" || response == "y" {
				break
			} else if response == "N" || response == "n" {
				os.Exit(0)
			}
		}
	}
	start := time.Now()
	result := fib(*nPtr)
	duration := time.Since(start)

	fmt.Printf("Fibonacci(%d) = %d\nExecution time: %s", *nPtr, result, duration)
}
func fib(n int) *big.Int {
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
