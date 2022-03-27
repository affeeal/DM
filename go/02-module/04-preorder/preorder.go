package main

import "fmt"

type Vertex struct {

	mark int
	next []*Vertex
}

func InitVertex(N int) *Vertex {

	var v Vertex
	v.mark = -1
	v.next = make([]*Vertex, N)

	for i := 0; i < N; i++ {

		v.next[i] = nil
	}
	return &v
}

func VisitVertex(V, N int, vs []*Vertex) {

	v := vs[V]
	v.mark = 0
	fmt.Printf("%d ", V)

	for i := 0; i < N; i++ {

		if v.next[i] != nil && v.next[i].mark == -1 {

			VisitVertex(i, N, vs)
		} 
	}
}

func main() {
	
	var N, M, V int
	fmt.Scanf("%d\n%d\n%d", &N, &M, &V)

	var vs []*Vertex = make([]*Vertex, N)
	for i := 0; i < N; i++ {

		vs[i] = InitVertex(N);
	}

	var u, v int
	for i := 0; i < M; i++ {

		fmt.Scanf("%d %d\n", &u, &v)
		vs[u].next[v] = vs[v]
		vs[v].next[u] = vs[u]
	}
	
	VisitVertex(V, N, vs);
	fmt.Println()
}