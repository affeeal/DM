package main

import "fmt"

type Vertex struct {

	x int
	l *List
}

type List struct {

	v *Vertex
	a byte
	next *List
}

func InitVertex(x int) *Vertex {

	v := new(Vertex)
	v.x = x
	v.l = InitList()
	return v
}

func InitList() *List {

	l := new(List)
	l.v = nil
	l.a = 0
	l.next = nil
	return l
}

func InitRoute(V int, vs []*Vertex) *List {

	r := InitList()
	r.next = InitList()
	r.next.v = vs[V]
	return r
}

func AddEdgeToList(u, v *Vertex, a byte, N int) {

	l := u.l
	for l.next != nil {

		if l.next.v.x > v.x {

			t := l.next
			l.next = InitList()
			l.next.v = v
			l.next.a = a
			l.next.next = t
			return
		}
		l = l.next
	}

	l.next = InitList()
	l.next.v = v
	l.next.a = a
}

func AddEdgeToRoute(u *Vertex, a byte, r *List) {

	l := r.next
	for l.next != nil {

		l = l.next
	}
	l.a = a
	l.next = InitList()
	l.next.v = u
}

func isEdgeInRoute(u, v *Vertex, a byte, r *List) bool {

	l := r.next
	for l.next != nil {

		if l.v == u && l.next.v == v || l.v == v && l.next.v == u {

			if l.a == a {

				return true
			}
		}
		l = l.next
	}
	return false
}

func ReturnToPrevVertex(r *List) {

	l := r.next
	prevl := l
	for l.next != nil {

		prevl = l
		l = l.next
	}
	l.a = prevl.a
	l.next = InitList()
	l.next.v = prevl.v
}

func BuildRoute(u *Vertex, N int, r *List) {

	hasFreeEdge := false

	for l := u.l.next; l != nil; l = l.next {

		if l.v == u && !isEdgeInRoute(u, l.v, l.a, r) {

			AddEdgeToRoute(l.v, l.a, r)
		} else if l.v.x > u.x {

			break
		}
	}
	
	for l := u.l.next; l != nil; l = l.next {

		if l.v != u && !isEdgeInRoute(u, l.v, l.a, r) {

			hasFreeEdge = true
			AddEdgeToRoute(l.v, l.a, r)
			BuildRoute(l.v, N, r)
		}
		
	}

	if !hasFreeEdge {

		ReturnToPrevVertex(r)
	}
}

func CorrectRoute(r *List) {

	l := r.next
	for l.next.next != nil {

		l = l.next
	}
	l.v = nil
	l.a = 0
	l.next = nil
}

func PrintRoute(r *List) {

	l := r.next
	for l.next.next != nil {

		fmt.Printf("%c ", l.a)
		l = l.next
	}
	fmt.Printf("%c\n", l.a)

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

func main() {
	
	var N, M, V int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)
	fmt.Scanf("%d", &V)

	vs := make([]*Vertex, N)
	for i := 0; i < N; i++ {

		vs[i] = InitVertex(i)
	}

	var u, v int
	var a byte
	for i := 0; i < M; i++ {

		fmt.Scanf("%d %d %c\n", &u, &v, &a)
		AddEdgeToList(vs[u], vs[v], a, N)
		if u != v {

			AddEdgeToList(vs[v], vs[u], a, N)
		}
	}

	r := InitRoute(V, vs)
	BuildRoute(vs[V], N, r)
	CorrectRoute(r)
	PrintRoute(r)
}