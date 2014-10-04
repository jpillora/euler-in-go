package main

import (
	"fmt"

	"github.com/jpillora/go-euler/euler"
)

var primes = make([]int, 1e3)

const (
	seqLength  = 4
	numFactors = 4
)

func main() {
	euler.FillWithPrimes(primes)

	var seq []int

	for n := 1; n < 9e9; n++ {
		//prime factor 'n' then store
		if factors(n, 0, []int{}) {
			seq = append(seq, n)
			if len(seq) == seqLength {
				fmt.Printf("%+v\n", seq)
				return
			}
		} else {
			seq = nil
		}
	}
}

func factors(n, u int, f []int) bool {
	//got all factors!
	if u == numFactors {
		return n == 1
	}
	//cant factor anymore!
	// if n == 0 {
	// 	return false
	// }

	for _, p := range primes {
		if n >= p && n%p == 0 && factors(n/p, incunqiue(u, p, f), append(f, p)) {
			return true
		}
	}
	return false
}

func incunqiue(u, n int, fs []int) int {
	for _, f := range fs {
		if f == n {
			return u
		}
	}
	return u + 1
}
