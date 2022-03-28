package main

import "fmt"

/*
 Доработать: подкорректировать приоритет выбора нового ребра:
 если возможно пройти по петле, то в первую очередь нужно сделать именно это.
*/

type Vertex struct {

	x int
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
	v.l = InitList()
	return v
}

func InitRoute(V int, vs []*Vertex) *List {

	r := InitList()
	r.next = InitList()
	r.next.v = vs[V]
	return r
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

func AddEdgeToRoute(u *Vertex, a byte, r *List) {

	l := r.next
	for l.next != nil {

		l = l.next
	}
	l.a = a
	l.next = InitList()
	l.next.v = u
}

func StepBackInRoute(r *List) {

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
	l := u.l.next
	for l != nil {

		if !isEdgeInRoute(u, l.v, l.a, r) {

			hasFreeEdge = true
			//fmt.Printf("GO FROM %d to %d by %c\n", u.x, l.v.x, l.a)
			//fmt.Printf("IS EDGE IN ROOT: %d\n", isEdgeInRoute(u, l.v, l.a, r))
			//PrintRoute(r)
			AddEdgeToRoute(l.v, l.a, r)
			BuildRoute(l.v, N, r)
		}
		l = l.next
	}

	if !hasFreeEdge {

		//fmt.Printf("STUCK IN %d\n", u.x)
		StepBackInRoute(r)
	}
}

func PrintRoute(r *List) {

	l := r.next
	for l.next.next != nil {

		fmt.Printf("%c ", l.a)
		l = l.next
	}
	//fmt.Printf("not ok")
	fmt.Printf("%c\n", l.a)

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
		AddEdge(u, v, N, a, vs)
		if u != v {

			AddEdge(v, u, N, a, vs)
		}
	}

	r := InitRoute(V, vs)
	//PrintGraph(N, vs)
	BuildRoute(vs[V], N, r)
	CorrectRoute(r)
	PrintRoute(r)
}