package euler

import (
	"sort"
	"strconv"
)

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

var globalsortcache = map[int]string{}

func SortDigits(n int) string {
	return SortDigitsWithCache(n, globalsortcache)
}

func SortDigitsWithCache(n int, sortcache map[int]string) string {
	var s string
	s, ok := sortcache[n]
	if ok {
		return s
	}
	s = strconv.Itoa(n)
	r := []rune(s)
	for len(r) < 4 {
		lenz := 4 - len(r)
		z := make([]rune, lenz)
		for i := 0; i < lenz; i++ {
			z[i] = '0'
		}
		r = append(z, r...)
	}
	sort.Sort(sortRunes(r))
	s = string(r)
	sortcache[n] = s
	return s
}
