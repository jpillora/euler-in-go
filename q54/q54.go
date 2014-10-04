package main

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strconv"
	"strings"
)

//royals
const (
	T = 10 << iota
	J
	Q
	K
	A
)

type rule func(h hand) []card

var rules = []rule{
	//royal flush
	func(h hand) []card {
		s := h[0].s
		v := A
		for i, c := range h {
			if c.v != v-i || c.s != s {
				return nil
			}
		}
		return []card{h[0]}
	},
	//straightflush
	func(h hand) []card {
		s := h[0].s
		v := h[0].v
		for i, c := range h {
			if c.v != v-i || c.s != s {
				return nil
			}
		}
		return []card{h[0]}
	},
	// four
	func(h hand) []card {
		t := 0
		if h[0].v != h[1].v {
			t = 1
		}
		v := h[t].v
		for i := 1; i <= 4; i++ {
			if h[t+i].v != v {
				return nil
			}
		}
		return []card{h[t]}
	},
	// full
	func(h hand) []card {
		return nil
	},
	// flush
	func(h hand) []card {
		return nil
	},
	// straight
	func(h hand) []card {
		return nil
	},
	// three
	func(h hand) []card {
		return nil
	},
	// twotwo
	func(h hand) []card {
		return nil
	},
	// two
	func(h hand) []card {
		var v int
		for i := 0; i < 5; i++ {
			if h[i].v == v {
				return []card{h[i]}
			}
			v = h[i].v
		}
		return nil
	},
	// high,
	func(h hand) []card {
		return h
	},
}

type cardstruct struct {
	v int
	s string
}

type card *cardstruct

type player struct {
	id, wins int
	hand     hand
}

func (p *player) newhand(cards []string) {
	hand := make([]card, len(cards))
	for i, c := range cards {
		vs := strings.Split(c, "")
		var v int
		switch vs[0] {
		case "A":
			v = A
		case "K":
			v = K
		case "Q":
			v = Q
		case "J":
			v = J
		case "T":
			v = T
		default:
			v, _ = strconv.Atoi(vs[0])
		}
		hand[i] = &cardstruct{v, vs[1]}
	}
	sort.Sort(byVal(hand))
	p.hand = hand
}

type hand []card

type bySuit hand

func (s bySuit) Less(i, j int) bool {
	return s[i].s > s[j].s
}

func (s hand) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s bySuit) Len() int {
	return len(s)
}

type byVal hand

func (s byVal) Less(i, j int) bool {
	return s[i].v > s[j].v
}

func (s byVal) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s byVal) Len() int {
	return len(s)
}

func compare(c1, c2 card) card {
	if c1.v > c2.v {
		return c1
	} else if c1.v < c2.v {
		return c2
	}
	return nil
}

func exec(r rule, p1, p2 *player) *player {
	set1 := r(p1.hand)
	set2 := r(p2.hand)
	if set1 != nil && set2 != nil {
		// fmt.Printf("#%d p1 p2 tie\n", n)
		//check highcards
		//will never tie highcards
		l := len(set1)
		if l != len(set2) {
			panic("invalid state!")
		}
		//compare player tie sets
		for i := 0; i < l; i++ {
			c1, c2 := set1[i], set2[i]
			c := compare(c1, c2)
			if c == c1 {
				return p1
			} else if c == c2 {
				return p2
			}
		}
	} else if set1 != nil {
		//p1 wins
		return p1
	} else if set2 != nil {
		//p2 wins
		return p1
	}
	//tie
	return nil
}

func main() {
	//create two players
	p1, p2 := &player{id: 1}, &player{id: 2}
	//read game
	b, err := ioutil.ReadFile("./poker.txt")
	if err != nil {
		panic(err)
	}
	//split game into rounds
	rounds := strings.Split(string(b), "\n")

	//split rounds into cards
	for _, r := range rounds {
		cards := strings.Fields(r)
		//provide each player with their hand
		p1.newhand(cards[:5])
		p2.newhand(cards[5:])
		//play!
		var winner *player
		for _, r := range rules {
			winner = exec(r, p1, p2)
			if winner != nil {
				break
			}
		}
		//tally winners
		if winner == p1 {
			p1.wins++
		} else if winner == p2 {
			p2.wins++
		}
	}
	fmt.Printf("p1 wins: %d, p2 wins: %d\n", p1.wins, p2.wins)

}
