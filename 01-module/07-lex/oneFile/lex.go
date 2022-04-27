package main

import (
	"fmt"
	"math/rand"
)

// AVLTree

var root *AVLTreeNode = nil

type AVLTreeNode struct {

	s string
	x, balance int
	left, right *AVLTreeNode
	parent *AVLTreeNode
}

func InitAVLTree() *AVLTreeNode {

	tree := AVLTreeNode { "", 0, 0, nil, nil, nil}
	root = &tree
	return &tree
}

func ReplaceNode(x, y *AVLTreeNode) {

	if x == root {

		root = y
		if y != nil {

			y.parent = nil
		}
	} else {

		p := x.parent
		if y != nil {

			y.parent = p
		}
		if p.left == x {

			p.left = y
		} else {

			p.right = y
		}
	}
}

func RotateLeft(x *AVLTreeNode) {

	y := x.right
	ReplaceNode(x, y)
	b := y.left
	if b != nil {

		b.parent = x
	}
	x.right = b
	x.parent = y
	y.left = x
	x.balance--
	if y.balance > 0 {

		x.balance -= y.balance
	}
	y.balance--
	if x.balance < 0 {

		y.balance += x.balance
	}
}

func RotateRight(x *AVLTreeNode) {

	y := x.left
	ReplaceNode(x, y)
	b := y.right
	if b != nil {

		b.parent = x
	}
	x.left = b
	x.parent = y
	y.right = x
	x.balance++
	if y.balance < 0 {

		x.balance -= y.balance
	}
	y.balance++
	if x.balance > 0 {

		y.balance += x.balance
	}
}

func (tree *AVLTreeNode) Lookup(s string) (x int, exists bool) {

	n := root
	for n != nil && n.s != s {

		if compareStrings(s, n.s) < 0 {

			n = n.left
		} else {

			n = n.right
		}
	}
	if n == nil {

		return 0, false
	}
	return n.x, true
}

func (tree *AVLTreeNode) Assign(s string, x int) {

	n := AVLTreeNode { s, x, 0, nil, nil, nil}
	if root.s == "" {

		root = &n
	} else {

		p := root
		for {

			if compareStrings(s, p.s) < 0 {

				if p.left == nil {

					p.left = &n
					n.parent = p
					break
				}
				p = p.left
			} else {

				if p.right == nil {

					p.right = &n
					n.parent = p
					break
				}
				p = p.right
			}
		}
	}
	pn := &n
	for {

		p := pn.parent
		if p == nil {

			break
		}
		if pn == p.left {

			p.balance--
			if p.balance == 0 {

				break
			}
			if p.balance == -2 {

				if pn.balance == 1 {

					RotateLeft(pn)
				}
				RotateRight(p)
				break
			}
		} else {

			p.balance++
			if p.balance == 0 {

				break
			}
			if p.balance == 2 {

				if pn.balance == -1 {

					RotateRight(pn)
				}
				RotateLeft(p)
				break
			}			
		}
		pn = p
	}
}

// SkipList

var m int = 8

type SkipList struct {

	s string
	x int
	next []*SkipList
}

func InitSkipList() *SkipList {

	l := SkipList { "", 0, make([]*SkipList, m) }
	for i := 0; i < m; i++ {

		l.next[i] = nil
	}
	return &l
}

func Succ(e *SkipList) *SkipList {

	return e.next[0]
}

func Skip(p []*SkipList, l *SkipList, s string) {

	e := l
	i := m - 1
	for i >= 0 {

		for e.next[i] != nil && compareStrings(e.next[i].s, s) < 0 {

			e = e.next[i]
		}
		p[i] = e
		i--
	}
}

func (l *SkipList) Lookup(s string) (x int, exists bool) {

	p := make([]*SkipList, m)
	Skip(p, l, s)
	e := Succ(p[0])
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

// main

type AssocArray interface {

	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
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

func lex(sentence string, array AssocArray) []int {

	var (
		begin, end int = -1, -1
		inWord bool = false
		x, sLen int = 1, len(sentence)
		result []int = []int {}
	)

	addToAssocArray := func() {

		var s string = string(sentence[begin:end])
		value, exists := array.Lookup(s)
		if exists {

			result = append(result, value)
		} else {

			array.Assign(s, x)
			result = append(result, x)
			x++
		}
	}

	for pos, sym := range sentence {

		if sym != ' ' && inWord == false {

			inWord = true
			begin = pos
		} else if sym == ' ' && inWord == true {

			inWord = false
			end = pos 
			addToAssocArray()
		}
	}
	if begin > end {

		end = sLen
		addToAssocArray()
	}
	return result
}

func main() {

	var (
		sentence string = "alpha x1 beta alpha x1 y "

		t *AVLTreeNode = InitAVLTree()
		arrayAVLTree AssocArray = t
		resultAVLTree []int = lex(sentence, arrayAVLTree)

		l *SkipList = InitSkipList()
		arraySL AssocArray = l
		resultSL = lex(sentence, arraySL)
	)

	fmt.Printf("SkipList: %d, AVLTree: %d.\n", resultSL, resultAVLTree)
}