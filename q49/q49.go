package main

import (
	"runtime"

	"github.com/jpillora/go-euler/euler"
)

func main() {
	//
	cpus := 1 // runtime.NumCPU()
	runtime.GOMAXPROCS(cpus)
	queue := make(chan int, cpus)
	//create a worker for each cpu
	for i := 0; i < cpus; i++ {
		w := &worker{
			queue:      queue,
			sortcache:  map[int]string{},
			primecache: map[int]bool{},
		}
		go w.dequeue()
	}
	//start assigning work
	for n := int(1e3); n < 10e3; n++ {
		queue <- n
	}
	//all done!
	close(queue)
}

type worker struct {
	queue      chan int
	sortcache  map[int]string
	primecache map[int]bool
}

func (w *worker) dequeue() {
	for {
		n, open := <-w.queue
		if !open {
			return
		}
		w.check(n)
	}
}

func (w *worker) check(n int) {
	if !w.prime(n) {
		return
	}
	for i := 0; n+i*2 < 10e3; i++ {
		nn := n + i
		if !w.prime(nn) || !w.palindrome(n, nn) {
			continue
		}
		nnn := nn + i
		if !w.prime(nnn) || !w.palindrome(nn, nnn) {
			continue
		}
		println(n, nn, nnn)
	}
}

func (w *worker) palindrome(a int, b int) bool {
	if a == b {
		return false
	}
	if euler.SortDigitsWithCache(a, w.sortcache) != euler.SortDigitsWithCache(b, w.sortcache) {
		return false
	}
	return true
}

// var primecache = map[int]bool{}
func (w *worker) prime(n int) bool {
	b, ok := w.primecache[n]
	if ok {
		return b
	}
	if n != 2 && n%2 == 0 {
		w.primecache[n] = false
		return false
	}
	i := n - 1
	for i > 1 {
		if n%i == 0 {
			w.primecache[n] = false
			return false
		}
		i--
	}
	w.primecache[n] = true
	return true
}
