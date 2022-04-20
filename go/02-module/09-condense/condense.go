package main

import "fmt"

var (
	time, count = 1, 1
	comps       []*List
)

// used types

type Stack struct {
	data     []*Vertex
	cap, top int
}

type List struct {
	v    *Vertex
	next *List
}

type ArcList struct {
	x, y int
	next *ArcList
}

type Vertex struct {
	num           int
	list          *List
	comp, low, T1 int
}

// Stack functions

func InitStack(n int) *Stack {

	s := new(Stack)
	s.data = make([]*Vertex, n)
	s.cap, s.top = 0, 0
	return s
}

func (s *Stack) Push(v *Vertex) {

	s.data[s.top] = v
	s.top++
}

func (s *Stack) Pop() *Vertex {

	s.top--
	return s.data[s.top]
}

// List functions

func InitList() *List {

	l := new(List)
	l.v = nil
	l.next = nil
	return l
}

func InsertList(u *Vertex) {

	l := comps[count-1]
	var p *List = nil
	for l.next != nil {
		if u.num < l.v.num {

			var tp **List = nil
			t := l

			if p == nil {
				tp = &comps[count-1]
			} else {
				tp = &p.next
			}

			*tp = InitList()
			(*tp).v = u
			(*tp).next = t
			return
		}
		p = l
		l = l.next
	}
	l.v = u
	l.next = InitList()
}

// ArcList functions

func InitArcList() *ArcList {

	a := new(ArcList)
	a.x, a.y = 0, 0
	a.next = nil
	return a
}

func AssignArcList(a *ArcList, x, y int) {

	for a.next != nil {

		a = a.next
	}
	a.x = x
	a.y = y
	a.next = InitArcList()
}

func LookupArcList(a *ArcList, x, y int) bool {

	for a.next != nil {

		if a.x == x && a.y == y {

			return true
		}
		a = a.next
	}
	return false
}

// Vertex functions

func InitVertex(i int) *Vertex {

	v := new(Vertex)
	v.num = i
	v.list = InitList()
	v.comp, v.low, v.T1 = 0, 0, 0
	return v
}

func (v *Vertex) InsertVertex(u *Vertex) {

	l := v.list
	for l.next != nil {

		l = l.next
	}
	l.v = u
	l.next = InitList()
}

// Auxiliary functions

func PrintVertices(vs []*Vertex) {

	fmt.Println("Printing vertices...")
	for _, v := range vs {

		fmt.Printf("%d(%d):", v.num, v.comp)
		for l := v.list; l.next != nil; l = l.next {

			fmt.Printf(" %d;", l.v.num)
		}
		fmt.Println()
	}
}

func PrintComponents() {

	fmt.Println("Printing components...")
	for i, c := range comps {

		fmt.Printf("%d:", i+1)
		for l := c; l.next != nil; l = l.next {

			fmt.Printf(" %d;", l.v.num)
		}
		fmt.Println()
	}
}

// Tarjan functions

func Tarjan(vs []*Vertex, N int) {

	s := InitStack(N)
	comps = make([]*List, 0)
	for _, v := range vs {

		if v.T1 == 0 {

			VisitVertex_Tarjan(vs, v, s)
		}
	}
}

func VisitVertex_Tarjan(vs []*Vertex, v *Vertex, s *Stack) {

	v.T1, v.low = time, time
	time++
	s.Push(v)
	for l := v.list; l.next != nil; l = l.next {

		u := l.v
		if u.T1 == 0 {

			VisitVertex_Tarjan(vs, u, s)
		}
		if u.comp == 0 && v.low > u.low {

			v.low = u.low
		}
	}
	if v.T1 == v.low {
		comps = append(comps, InitList())
		for {
			u := s.Pop()
			u.comp = count
			InsertList(u)
			if u == v {
				break
			}
		}
		count++
	}
}

// Main functions

func PrintCondenseGraph(vs []*Vertex) {

	fmt.Println("digraph {")
	for i, c := range comps {

		fmt.Printf("\t%d [label=\"[", i+1)
		for l := c; l.next != nil; l = l.next {
			if l.next.next == nil {
				fmt.Printf("%d]\"]\n", l.v.num)
			} else {
				fmt.Printf("%d ", l.v.num)
			}
		}
	}
	a := InitArcList()
	for _, v := range vs {
		for l := v.list; l.next != nil; l = l.next {

			u := l.v
			vc, uc := v.comp, u.comp
			if vc != uc && !LookupArcList(a, vc, uc) {

				AssignArcList(a, vc, uc)
				fmt.Printf("\t%d->%d\n", vc, uc)
			}
		}

	}
	fmt.Println("}")
}

func main() {

	var N, M int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)

	vs := make([]*Vertex, N)
	for i := 0; i < N; i++ {

		vs[i] = InitVertex(i)
	}

	var v, u int
	for i := 0; i < M; i++ {

		fmt.Scanf("%d %d\n", &v, &u)
		vs[v].InsertVertex(vs[u])
	}

	Tarjan(vs, N)
	PrintCondenseGraph(vs)
}
