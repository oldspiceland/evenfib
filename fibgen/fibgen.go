package fibgen

import (
	"fmt"
)

func findNext(fib1, fib2 int) (fib3 int) {
	fib3 = fib1 + fib2
	return
}

func GenerateSeq(seed1, seed2, limit int) {
	fib := 0
	for fib < limit {
		fib = seed1 + seed2
		fmt.Printf("The next number in the sequence is %v\n",fib)
		seed1 = seed2
		seed2 = fib
	}
}
