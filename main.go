package main

import (
	"fmt"
	"github.com/oldspiceland/evenfib/fibgen"
)

func main() {
	seq := fibgen.GenerateSeq(0, 1, 4000000)

	fibSum := summer(seq)
	fmt.Printf("The sum of the even sequence members is %v\n", fibSum)
}

func summer(seq []int) (sum int) {
	sum = 0
	for i := 0; i < len(seq); i++ {
		if seq[i]%2 == 0 {
			sum += seq[i]
		}
	}
	return
}
