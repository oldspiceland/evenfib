package main

import (
	"flag"
	"fmt"
	"github.com/oldspiceland/evenfib/fibgen"
)

func main() {
	seedOne := 0
	seedTwo := 1
	var flagSeed = flag.Bool("seed", false, "Flag for user indicating she wants to change starting sequence member")

	flag.Parse()

	if *flagSeed == true {
		seedOne, seedTwo = getSeed()
	}
	seq := fibgen.GenerateSeq(seedOne, seedTwo, 4000000)

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

func getSeed() (seedOne int, seedTwo int) {
	seed := 0
	fmt.Print("Enter first sequence member: ")
	_, _ = fmt.Scanf("%v", &seed)
	seedOne = seed
	seedTwo = seedOne + seedOne
	return
}
