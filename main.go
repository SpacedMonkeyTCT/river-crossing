package main

import "fmt"

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

	fmt.Printf("left : %v\n", l)
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

func main() {
	history := []state{{left, left, left, left}}

	history[0].display()
}
