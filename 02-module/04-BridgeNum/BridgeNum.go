package main

import (
	"container/list"
	"fmt"
)

type Vertex struct {
	x, mark, comp int
	parent        *Vertex
	l             *list.List
}

// x is needed for debug
// mark: -1 == white, 0 == gray, 1 == black

type Vertices []*Vertex

type Queue struct {
	data                   Vertices
	cap, count, head, tail int
}

// queue functions

func InitQueue(N int) *Queue {

	q := new(Queue)
	q.data = make(Vertices, N)
	q.cap = N
	q.count, q.head, q.tail = 0, 0, 0
	return q
}

func QueueEmpty(q *Queue) bool {

	return q.count == 0
}

func Enqueue(q *Queue, v *Vertex) {

	q.data[q.tail] = v
	q.tail++
	if q.tail == q.cap {
		q.tail = 0
	}
	q.count++
}

func Dequeue(q *Queue) *Vertex {

	v := q.data[q.head]
	q.head++
	if q.head == q.cap {
		q.head = 0
	}
	q.count--
	return v
}

// vertex functions

func InitVertex(i int) *Vertex {

	v := new(Vertex)
	v.x = i
	v.mark, v.comp = -1, -1
	v.parent = nil
	v.l = list.New()
	return v
}

// main functions

func DFS1(vs *Vertices, q *Queue) {

	for _, v := range *vs {
		if v.mark == -1 {
			VisitVertex1(vs, v, q)
		}
	}
}

func VisitVertex1(vs *Vertices, v *Vertex, q *Queue) {

	v.mark = 0
	Enqueue(q, v)
	for e := v.l.Front(); e != nil; e = e.Next() {

		u := e.Value.(*Vertex)
		if u.mark == -1 {
			u.parent = v
			VisitVertex1(vs, u, q)
		}
	}
	v.mark = 1
}

func DFS2(vs *Vertices, q *Queue) {

	component := 0
	for !QueueEmpty(q) {
		v := Dequeue(q)
		if v.comp == -1 {
			VisitVertex2(vs, v, component)
			component++
		}
	}
}

func VisitVertex2(vs *Vertices, v *Vertex, component int) {
	v.comp = component
	for e := v.l.Front(); e != nil; e = e.Next() {
		u := e.Value.(*Vertex)
		if u.comp == -1 && u.parent != v {
			VisitVertex2(vs, u, component)
		}
	}
}

// auxiliary functions

func PrintVertices(vs *Vertices) {

	for _, v := range *vs {
		fmt.Printf("%d(%d):", v.x, v.comp)
		for e := v.l.Front(); e != nil; e = e.Next() {
			fmt.Printf(" %d;", e.Value.(*Vertex).x)
		}
		fmt.Println()
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

	q := InitQueue(N)
	DFS1(&vs, q)
	DFS2(&vs, q)

	//PrintVertices(&vs)

	bridges := 0
	for _, v := range vs {
		for e := v.l.Front(); e != nil; e = e.Next() {

			u := e.Value.(*Vertex)
			if v.x < u.x && v.comp != u.comp {
				bridges++
			}
		}
	}

	fmt.Println(bridges)
}
