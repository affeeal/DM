package main

import "math/rand"

var m int = 8

type SkipList struct {
	s string
	x int
	next []*SkipList
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
		r int = rand.Intn(10) * 2
		i int = 0
	)
	Skip(p, l, s)
	e.next = make([]*SkipList, m)
	e.s = s
	e.x = x
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