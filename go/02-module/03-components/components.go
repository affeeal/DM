package main

import "fmt"

type Vertex struct {

	x, cmp int
	l *List
}

type Vertexes []*Vertex

type List struct {

	v *Vertex
	next *List
}

func InitVertex(i int) *Vertex {

	v := new(Vertex)
	v.x = i
	v.cmp = -1
	v.l = new(List)
	v.l.v = nil
	v.l.next = nil
	return v
}

func InitVertexes(N int) *Vertexes {

	vs := make(Vertexes, N)
	for i := range vs {

		vs[i] = InitVertex(i)
	}
	return &vs
}

func EndOfList(l *List) *List{

	for l.next != nil {

		l = l.next
	}
	return l
}

func AddEdge(v1, v2 *Vertex) {

	l1 := EndOfList(v1.l)
	l1.v = v2
	l1.next = new(List)
	l1.next.v = nil
	l1.next.next = nil

	l2 := EndOfList(v2.l)
	l2.v = v1
	l2.next = new(List)
	l2.next.v = nil
	l2.next.next = nil
}

func BuildCmp(v *Vertex, cmp int) {

	v.cmp = cmp
	l := v.l
	for l.next != nil {

		if l.v.cmp != cmp {

			BuildCmp(l.v, cmp)
		}
		l = l.next
	}
}

func CheckResult(N int, vs *Vertexes) {

	for i := 0; i < N; i++ {

		fmt.Printf("%d(%d) -> ", (*vs)[i].x, (*vs)[i].cmp)
		for l := (*vs)[i].l; l.next != nil; l = l.next {

			fmt.Printf("%d(%d) ", l.v.x, (*vs)[i].cmp)
		}
		fmt.Println()
	}
}

func main() {
	
	var N, M int
	fmt.Scanf("%d\n%d", &N, &M)
	vs := InitVertexes(N)
	var v1, v2 int
	for i := 0; i < M; i++ {  //  M - 1 ?

		fmt.Scanf("%d %d\n", &v1, &v2)
		AddEdge((*vs)[v1], (*vs)[v2])
	}
	cmp := 0
	for i := range (*vs) {

		if (*vs)[i].cmp == -1 {

			BuildCmp((*vs)[i], cmp)
			cmp++
		}
	}
	//CheckResult(N, vs)
	fmt.Printf("%d\n", cmp)
}