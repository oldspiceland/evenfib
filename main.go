package main

import (
	"github.com/oldspiceland/evenfib/fibgen"
)

func main() {
	seq := fibgen.GenerateSeq(1, 2, 4000000)

	fibSum := summer(seq)
	fmt.Printf("The sum of the even sequence members is %v\n", fibSum)
}

func summer(seq int) (sum int) {
	sum := 0
	for i = 0; i <= len(seq); i++ {
		if seq[i]%2 == 0 {
			sum += seq[i]
		}
	}
	return
}
