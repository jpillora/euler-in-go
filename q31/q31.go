package main

func main() {
	// 1p, 2p, 5p, 10p, 20p, 50p, £1 (100p) and £2 (200p).
	coins := []uint16{1, 2, 5, 10, 20, 50, 100, 200}
	n := 0
	var s func(uint16, uint16)
	s = func(last uint16, total uint16) {
		if total == 200 {
			n++
		} else if total > 200 {
			return
		}
		for _, c := range coins {
			if c >= last {
				s(c, total+c)
			}
		}
	}
	s(0, 0)
	print(n)
}
