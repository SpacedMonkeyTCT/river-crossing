package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"time"
)

type char int

const (
	farmer char = iota
	goose
	bean
	fox
	numChar
)

func (c char) String() string {
	charName := []string{"Farmer", "Goose", "Bean", "Fox"}
	return charName[c]
}

func (c char) eats(e char, s state) bool {
	if c == fox && e == goose || c == goose && e == bean {
		if s[c] == s[e] && s[farmer] != s[c] {
			return true
		}
	}
	return false
}

type place int

const (
	left place = iota
	boat
	right
	numPlace
)

func (p place) String() string {
	placeName := []string{"Left", "Boat", "Right"}
	return placeName[p]
}

type state [numChar]place

func (s state) display() {
	l := []char{}
	b := []char{}
	r := []char{}

	for c, p := range s {
		switch p {
		case left:
			l = append(l, char(c))
		case boat:
			b = append(b, char(c))
		case right:
			r = append(r, char(c))
		}
	}

	fmt.Printf("\nleft : %v\n", l)
	fmt.Printf("boat : %v\n", b)
	fmt.Printf("right: %v\n", r)
}

func (s state) destinations() []place {
	if s[farmer] == boat {
		return []place{left, right}
	}
	return []place{boat}
}

func (s state) travellers() []char {
	t := []char{}
	for c, p := range s {
		if p == s[farmer] {
			t = append(t, char(c))
		}
	}
	return t
}

func (s state) generateMoves() []state {
	d := s.destinations()
	t := s.travellers()

	ns := []state{}
	for _, p := range d {
		for _, c := range t {
			m := s
			m[farmer] = p
			m[c] = p
			ns = append(ns, m)
		}
	}

	return ns
}

func (s state) inHistory(history []state) bool {
	for _, h := range history {
		if reflect.DeepEqual(s, h) {
			return true
		}
	}
	return false
}

func (s state) boatSinks() bool {
	c := 0
	for _, p := range s {
		if p == boat {
			c++
		}
	}
	if c > 2 {
		return true
	}
	return false
}

func main() {
	s := state{left, left, left, left}
	end := state{right, right, right, right}
	badMoves := []state{}
	previousMoves := []state{s}
	rand.Seed(time.Now().Unix())

	for !reflect.DeepEqual(s, end) {
		pm := s.generateMoves()
		ns := []state{}

		for _, m := range pm {
			if fox.eats(goose, m) {
				continue
			}
			if goose.eats(bean, m) {
				continue
			}
			if m.boatSinks() {
				continue
			}
			if m.inHistory(badMoves) {
				continue
			}
			if m.inHistory(previousMoves) {
				continue
			}
			ns = append(ns, m)
		}
		// if there are no moves, we must have made a wrong choice in the past
		// mark our current state as bad and rollback history
		if len(ns) == 0 {
			badMoves = append(badMoves, s)
			previousMoves = previousMoves[:len(previousMoves)-1]
			s = previousMoves[len(previousMoves)-1]
		} else {
			s = ns[rand.Intn(len(ns))]
			s.display()
			previousMoves = append(previousMoves, s)
		}
	}

	fmt.Println()
	fmt.Println("SOLUTION FOUND!")

	for _, s := range previousMoves {
		s.display()
	}
}
