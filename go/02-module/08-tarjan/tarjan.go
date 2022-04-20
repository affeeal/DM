package main

import "fmt"

var (
	time, count = 1, 1
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

type Vertex struct {
	num           int //  needed for debug
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

	for _, v := range vs {

		fmt.Printf("%d(%d):", v.num, v.comp)
		for l := v.list; l.next != nil; l = l.next {

			fmt.Printf(" %d;", l.v.num)
		}
		fmt.Println()
	}
}

// Main functions

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
		for {
			u := s.Pop()
			u.comp = count
			if u == v {
				break
			}
		}
		count++
	}
}

func Tarjan(vs []*Vertex, N int) {

	s := InitStack(N)
	for _, v := range vs {

		if v.T1 == 0 {

			VisitVertex_Tarjan(vs, v, s)
		}
	}
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

	PrintVertices(vs)

	count--
	fmt.Println(count)
}
