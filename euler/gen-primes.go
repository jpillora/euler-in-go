package euler

func FillWithPrimes(slice []int) {
	g := PrimesGenerator()
	for i := 0; i < len(slice); i++ {
		slice[i] = <-g
	}
}

func PrimesGenerator() chan int {
	ch := make(chan int, 1)

	go func() {
		p := 0
		for {
			if IsPrime(p) {
				ch <- p
			}
			p++
		}
	}()

	return ch
}
