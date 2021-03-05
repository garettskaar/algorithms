package main

import (
	"flag"
	"fmt"
	"github/garettskaar/algorithms/fibonacci"
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
	result := fibonacci.FibDynamic(*nPtr)
	duration := time.Since(start)

	fmt.Printf("Fibonacci(%d) = %d\nExecution time: %s", *nPtr, result, duration)
}
