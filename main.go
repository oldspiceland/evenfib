package main

import (
	"flag"
	"fmt"
	"github.com/oldspiceland/evenfib/fibgen"
)

func main() {
	var flagLimit = flag.Int("limit", 1000, "Limit on the size of sequence members")

	flag.Parse()

	seq := fibgen.GenerateSeq(0, 1, *flagLimit)

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
