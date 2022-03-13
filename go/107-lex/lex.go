package main

import (
	"fmt"
	"math/rand"
)

var m int = 4

type SkipList struct {
	s string
	x int
	next []*SkipList
}

func compareStrings(a, b string) int {

	var (
		lenA int = len(a)
		lenB int = len(b)
	)
	if lenA < lenB {

		return -1
	} else if lenA > lenB {

		return 1
	} else {

		for i := 0; i < lenA; i++ {

			if a[i] > b[i] {

				return -1
			} else if a[i] < b[i] {

				return 1
			}
		}
		return 0
	}
}

func InitSkipList() *SkipList {

	var l SkipList
	l.next = make([]*SkipList, m)
	for i := 0; i < m; i++ {

		l.next[i] = nil
	}
	return &l
}

func Succ(e *SkipList) *SkipList {

	return e.next[0]
}

func Skip(p []*SkipList, l *SkipList, s string) {

	var (
		e *SkipList = l
		i int = m - 1
	)

	for i >= 0 {

		for e.next[i] != nil && compareStrings(e.next[i].s, s) < 0 {

			e = e.next[i]
		}
		p[i] = e
		i--
	}
}

func (l *SkipList) Lookup(s string) (x int, exists bool) {

	var (
		p []*SkipList = make([]*SkipList, m)
		e *SkipList
	)

	Skip(p, l, s)
	e = Succ(p[0])
	if e == nil || e.s != s {

		return 0, false
	}
	return e.x, true
}

func (l *SkipList) Assign(s string, x int) {

	var (
		p []*SkipList = make([]*SkipList, m)
		e SkipList
		r, i int = rand.Intn(10) * 2, 0
	)

	Skip(p, l, s)
	if p[0].next[0] != nil && p[0].next[0].s == s {

		fmt.Printf("something went wrong...\n")
		return
	}
	e.next = make([]*SkipList, m)
	for i < m && r % 2 == 0 {

		e.next[i] = p[i].next[i]
		p[i].next[i] = &e
		i++
		r /= 2
	}
	for i < m {

		e.next[i] = nil
		i++
	}
}

type AssocArray interface {

	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

func lex(sentence string, array AssocArray) []int {

	var (
		begin, end int = -1, -1
		inWord bool = false
		x, sLen int = 0, len(sentence)
		lex []int = []int {}
	)

	addToAssocArray := func() {

		s := string(sentence[begin:end])
		value, exists := array.Lookup(s)
		if exists {

			lex = append(lex, value)
		} else {

			array.Assign(s, x)
			lex = append(lex, x)
			x++
		}
	}

	for pos, sym := range sentence {

		if sym != ' ' && inWord == false {

			inWord = true
			begin = pos
		} else if sym == ' ' && inWord == true {

			inWord = false
			end = pos - 1 
			addToAssocArray()
		}
	}
	if begin > end {

		end = sLen - 1
		addToAssocArray()
	}
	return lex
}

func main() {

	var sentence string = "a a a"
	var (
		l *SkipList = InitSkipList()
		arraySL AssocArray = l
	)
	resultSL := lex(sentence, arraySL)
	fmt.Printf("%q\n", resultSL)
}