package main

import (
	"fmt"
)

type Vertex struct {

	x int
	p []*Vertex
	l *List
}

type List struct {

	v *Vertex
	a byte
	next *List
}

func InitList() *List {

	l := new(List)
	l.v = nil
	l.a = 0
	l.next = nil
	return l
}

func InitVertex(x int) *Vertex {

	v := new(Vertex)
	v.x = x
	v.p = make([]*Vertex, 0)
	v.l = InitList()
	return v
}

func AddEdge(u, v, N int, a byte, vs []*Vertex) {

	l := vs[u].l
	for l.next != nil {

		if l.next.v.x > v {

			t := l.next
			l.next = InitList()
			l.next.v = vs[v]
			l.next.a = a
			l.next.next = t
			return
		}
		l = l.next
	}

	l.next = InitList()
	l.next.v = vs[v]
	l.next.a = a
}

func PrintGraph(N int, vs []*Vertex) {

	for i := 0; i < N; i++ {

		v := vs[i]
		fmt.Printf("%d => ", v.x)
		for l := v.l.next; l != nil; l = l.next {

			fmt.Printf("%d, %c; ", l.v.x, l.a)
		}
		fmt.Println()
	}
}

func areVerticesAdjacent (u, v int, vs []*Vertex) (a byte, verdict bool) {

	if u < v {

		u, v = v, u
	}

	l := vs[u].l.next
	for l != nil {

		if l.v.x == v {

			return l.a, true
		} else if l.v.x > v {

			return 0, false
		}

		l = l.next
	}
	return 0, false
}

func isParent(u, v int, vs []*Vertex) bool { //  is "u" parent for "v"

	p := vs[v].p
	for i := range p {

		if p[i] == vs[u] {

			return true
		}
	}
	return false
}

func FindParent(u int, vs []*Vertex) (int, byte) {

	l := vs[u].l.next
	for l != nil {

		if isParent(l.v.x, u, vs) {

			return l.v.x, l.a
		}
		l = l.next
	}
	return 0, 0
}

func FindRoute(V, N int, vs []*Vertex) {

	for i := 0; i < N; i++ {

		a, verdict := areVerticesAdjacent(V, i, vs)
		if verdict && !(isParent(V, i, vs) || isParent(i, V, vs)) {

			vs[i].p = append(vs[i].p, vs[V])
			fmt.Printf("%c ", a)
			FindRoute(i, N, vs)
			return
		}
	}
	p, a := FindParent(V, vs)
	fmt.Printf("%c ", a)
	FindRoute(p, N, vs)
}

func main() {
	
	var N, M, V int
	fmt.Scanf("%d\n%d\n%d\n", &N, &M, &V)

	vs := make([]*Vertex, N)
	for i := 0; i < N; i++ {

		vs[i] = InitVertex(i)
	}

	var u, v int
	var a byte
	for i := 0; i < M; i++ {

		fmt.Scanf("%d %d %c\n", &u, &v, &a)
		AddEdge(u, v, N, a, vs)
		if u != v {

			AddEdge(v, u, N, a, vs)
		}
	}

	FindRoute(V, N, vs)
}