package main

import (
	"container/list"
	"fmt"
)

type Vertex struct {
	x int
	l *list.List
}

type PivotVertex struct {
	v    *Vertex
	dist []int
}

type Vertices []*Vertex
type PivotVertices []*PivotVertex

// vertex functions

func InitVertex(i int) *Vertex {

	v := new(Vertex)
	v.x = i
	v.l = list.New()
	return v
}

// PivotVertex functions

func InitPivotVertex(v *Vertex, N int) *PivotVertex {
	pv := new(PivotVertex)
	pv.v = v
	pv.dist = make([]int, N)
	for i := 0; i < N; i++ {
		pv.dist[i] = 0
	}
	return pv
}

// auxiliary functions

func PrintVertices(vs *Vertices) {

	for _, v := range *vs {
		fmt.Printf("%d:", v.x)
		for e := v.l.Front(); e != nil; e = e.Next() {
			fmt.Printf(" %d;", e.Value.(*Vertex).x)
		}
		fmt.Println()
	}
}

func PrintPivotVertices(pvs *PivotVertices) {

	for _, pv := range *pvs {
		fmt.Printf("%d: [", pv.v.x)
		for _, d := range pv.dist {
			fmt.Printf(" %d", d)
		}
		fmt.Println(" ]")
	}
}

// main functions

func FindDistances(pv *PivotVertex, v *Vertex, d int) {

	for e := v.l.Front(); e != nil; e = e.Next() {

		u := e.Value.(*Vertex)
		if pv.dist[u.x] == 0 && u.x != pv.v.x || pv.dist[u.x] > d {
			pv.dist[u.x] = d
		}
	}
	for e := v.l.Front(); e != nil; e = e.Next() {

		u := e.Value.(*Vertex)
		if pv.dist[u.x] == d {
			FindDistances(pv, u, d+1)
		}
	}
}

// main function

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
		vs[u].l.PushBack(vs[v])
	}

	var K int
	fmt.Scanf("%d\n", &K)

	var k int
	pvs := make(PivotVertices, K)
	for i := 0; i < K; i++ {
		fmt.Scanf("%d", &k)
		pvs[i] = InitPivotVertex(vs[k], N)
	}

	//PrintVertices(&vs)

	for _, pv := range pvs {
		FindDistances(pv, pv.v, 1)
	}

	//PrintPivotVertices(&pvs)

	hasEqualDist := false
	for i := 0; i < N; i++ {

		d := pvs[0].dist[i]
		isEqualDist := true
		for j := 1; j < K; j++ {

			if pvs[j].dist[i] != d {
				isEqualDist = false
				break
			}
		}
		if d != 0 && isEqualDist {
			hasEqualDist = true
			fmt.Println(vs[i].x)
		}
	}
	if !hasEqualDist {
		fmt.Println("-")
	}
}
