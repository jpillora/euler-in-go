package euler

var primecache = map[int]bool{}

func IsPrime(n int) bool {
	if n == 2 {
		return true
	}
	if n == 1 || n%2 == 0 {
		return false
	}

	b, ok := primecache[n]
	if ok {
		return b
	}
	i := n - 1
	for i > 1 {
		if n%i == 0 {
			primecache[n] = false
			return false
		}
		i--
	}
	primecache[n] = true
	return true
}
