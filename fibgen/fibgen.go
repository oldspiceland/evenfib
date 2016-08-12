package fibgen

import (
	"fmt"
)

func findNext(fib1, fib2 int) (fib3 int) {
	fib3 = fib1 + fib2
	return
}

func GenerateSeq(seed1, seed2, limit int) int { //GenerateSeq takes two seeds and a limit
	seq := make([]int, 10000)
	fib := 0
	for fib < limit {
		fib = seed1 + seed2
		if fib < limit {
			fmt.Printf("The next number in the sequence is %v\n", fib)
			seed1 = seed2
			seed2 = fib
			seq = append(seq, fib)
		}
	}
	return seq
}
