package main

import (
	"fmt"
	"os"

	"github.com/jpillora/go-euler/euler"
)

const max = int(1e6)

//contains the current length and value of the sum
type sum struct {
	length, val int
}

func main() {
	var long sum
	sums := []*sum{}
	for p := 1; p < max; p++ {
		//p is the next prime in sequence
		if !euler.IsPrime(p) {
			continue
		}
		//add p to all sums
		for n := 0; n < len(sums); {
			s := sums[n]
			s.length++
			s.val += p
			//sum is over max - remove it!
			if s.val >= max || s.val+(long.length-s.length)*p >= max {
				sums = append(sums[:n], sums[n+1:]...)
				//removed all! no more primes are eligible
				if len(sums) == 0 {
					fmt.Printf("%+v\n", long)
					os.Exit(0)
				}
				continue
			}
			//track closest to max
			if s.val > long.val && s.length > long.length && euler.IsPrime(s.val) {
				long = *s
			}
			n++
		}
		//add a new prime to the sums list
		sums = append(sums, &sum{length: 1, val: p})
	}
}
