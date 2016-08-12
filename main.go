package main

import (
	"flag"
	"fmt"
	"github.com/oldspiceland/evenfib/fibgen"
)

func main() {
	var flagLimit = flag.Int("limit", 1000, "Limit on the size of sequence members")

	seedOne := 0
	seedTwo := 1
	var flagSeed = flag.Bool("seed", false, "Flag for user indicating she wants to change starting sequence member")

	flag.Parse()

	if *flagSeed == true {
		seedOne, seedTwo = getSeed()
	}
	seq := fibgen.GenerateSeq(seedOne, seedTwo, *flagLimit)

	fibSum := summer(seq)
	fmt.Printf("The sum of the even sequence members is %v\n", fibSum)
}

//summer takes a slice of int, returns an int
func summer(seq []int) (sum int) {
	sum = 0
	for i := 0; i < len(seq); i++ {
		if seq[i]%2 == 0 {
			sum += seq[i]
		}
	}
	return
}

//getSeed takes two no args, returns two ints, gets one user input
func getSeed() (seedOne int, seedTwo int) {
	seed := 0
	fmt.Print("Enter first sequence member: ")
	_, _ = fmt.Scanf("%v", &seed)
	seedOne = seed
	seedTwo = seedOne + seedOne
	return
}
