package main

import "github.com/jpillora/go-euler/euler"

const (
	conseq = 6
)

func main() {
	for i := 1; ; i++ {
		all := true
		d := euler.SortDigits(i)
		for c := 2; c <= conseq; c++ {
			if d != euler.SortDigits(i*c) {
				all = false
				break
			}
		}
		if all {
			println(i)
			return
		}
	}
}
