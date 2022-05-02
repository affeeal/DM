package main

import (
	"container/list"
	"fmt"
)

var (
	time  = 1
	count = 1
)

type Stack struct {
	data     Vertices
	cap, top int
}

type Vertex struct {
	num           int
	comp, T1, low int
	l             *list.List
}

type Vertices []*Vertex

type Comp struct {
	vs      *list.List
	parents *list.List
}

type Comps []*Comp

// Vertex functions

func InitVertex(i int) *Vertex {
	v := new(Vertex)
	v.num = i
	v.comp, v.low, v.T1 = 0, 0, 0
	v.l = list.New()
	return v
}

// Comp functions

func InitComp() *Comp {
	c := new(Comp)
	c.vs = list.New()
	c.parents = list.New()
	return c
}

// Stack functions

func InitStack(N int) *Stack {

	s := new(Stack)
	s.data = make(Vertices, N)
	s.cap, s.top = 0, 0
	return s
}

func Push(s *Stack, v *Vertex) {

	s.data[s.top] = v
	s.top++
}

func Pop(s *Stack) *Vertex {

	s.top--
	return s.data[s.top]
}

// auxiliary functions

func PrintComps(cs *Comps) {
	for i, c := range *cs {
		fmt.Printf("%d:", i+1)
		for e := c.vs.Front(); e != nil; e = e.Next() {
			v := e.Value.(*Vertex)
			fmt.Printf(" %d;", v.num)
		}
		fmt.Println()
		if c.parents.Front() != nil {
			fmt.Print("Inherited by")
			for e := c.parents.Front(); e != nil; e = e.Next() {
				fmt.Printf(" %d;", e.Value.(*Comp).vs.Front().Value.(*Vertex).comp)
			}
			fmt.Println()
		}
	}
}

func PrintVertices(vs *Vertices) {
	for _, v := range *vs {
		fmt.Printf("%d(%d):", v.num, v.comp)
		for e := v.l.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			fmt.Printf(" %d;", u.num)
		}
		fmt.Println()
	}
}

// main

func Tarjan(vs *Vertices, cs *Comps) {
	s := InitStack(len(*vs))
	for _, v := range *vs {
		if v.T1 == 0 {
			VisitVertexTarjan(vs, v, s, cs)
		}
	}
}

func VisitVertexTarjan(vs *Vertices, v *Vertex, s *Stack, cs *Comps) {
	v.T1, v.low = time, time
	time++
	Push(s, v)
	for e := v.l.Front(); e != nil; e = e.Next() {
		u := e.Value.(*Vertex)
		if u.T1 == 0 {
			VisitVertexTarjan(vs, u, s, cs)
		}
		if u.comp == 0 && v.low > u.low {
			v.low = u.low
		}
	}
	if v.T1 == v.low {
		*cs = append(*cs, InitComp())
		for {
			u := Pop(s)
			u.comp = count
			(*cs)[count-1].vs.PushFront(u)
			if u == v {
				break
			}
		}
		count++
	}
}

func SetRelations(cs *Comps) {

	for _, c := range *cs {
		for e := c.vs.Front(); e != nil; e = e.Next() {
			v := e.Value.(*Vertex)
			for e := v.l.Front(); e != nil; e = e.Next() {
				u := e.Value.(*Vertex)
				if v.comp != u.comp {
					isParent := false
					for e := (*cs)[u.comp-1].parents.Front(); e != nil; e = e.Next() {
						parent := e.Value.(*Comp)
						if c == parent {
							isParent = true
							break
						}
					}
					if !isParent {
						(*cs)[u.comp-1].parents.PushBack(c)
					}
				}
			}

		}
	}
}

func main() {

	var N, M int
	fmt.Scanf("%d\n", &N)
	fmt.Scanf("%d\n", &M)

	vs := make(Vertices, N)
	for i := 0; i < N; i++ {
		vs[i] = InitVertex(i)
	}

	var v, u int
	for i := 0; i < M; i++ {
		fmt.Scanf("%d %d\n", &v, &u)
		vs[v].l.PushBack(vs[u])
	}

	cs := make(Comps, 0)
	Tarjan(&vs, &cs)
	SetRelations(&cs)

	for _, c := range cs {
		if c.parents.Front() == nil {
			fmt.Printf("%d\n", c.vs.Front().Value.(*Vertex).num)
		}
	}
}
