package main

import (
	"fmt"
	"math"
)

func main() {

	pent := func(n int) int {
		return n * ((3 * n) - 1) / 2
	}

	//inverse pent
	ispent := func(p int) bool {
		f := (math.Sqrt(1+24*float64(p)) + 1.0) / 6.0
		return f == math.Floor(f)
	}

	//try all j and k up to 10,000
	max := int(1e4)
	for j := 1; j < max; j++ {
		for k := j; k < max; k++ {
			pj := pent(j)
			pk := pent(k)
			sum := pj + pk
			dif := pk - pj
			if ispent(sum) && ispent(dif) {
				fmt.Printf("%d,%d -> sum: %d  diff: %d\n", j, k, sum, dif)
				return
			}
		}
	}
}
