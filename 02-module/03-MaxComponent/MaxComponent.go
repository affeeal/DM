package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
)

type Vertex struct {
	x, mark, comp int
	parent        *Vertex
	l             *list.List
}

// mark: -1 == white, 0 == gray, 1 == black

type Vertices []*Vertex

type Component struct {
	num, vCount int
	eCount      float32
	repr        *Vertex
}

// num is a number of a component
// vCount is a number of vertices in the component
// eCount is a number of edges in the component
// repr is a representative of a component

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
	(*q).head++
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

// component functions

func InitComponent(i int) *Component {
	c := new(Component)
	c.num = i
	c.vCount, c.eCount = 0, 0.0
	c.repr = nil
	return c
}

// main functions

func DFS1(vs Vertices, q *Queue) {
	for _, v := range vs {
		if v.mark == -1 {
			VisitVertex1(vs, v, q)
		}
	}
}

func VisitVertex1(vs Vertices, v *Vertex, q *Queue) {
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

func DFS2(vs Vertices, q *Queue) int {
	component := 0
	for !QueueEmpty(q) {
		v := Dequeue(q)
		if v.comp == -1 {
			VisitVertex2(vs, v, component)
			component++
		}
	}
	return component
}

func VisitVertex2(vs Vertices, v *Vertex, component int) {
	v.comp = component
	for e := v.l.Front(); e != nil; e = e.Next() {
		u := e.Value.(*Vertex)
		if u.comp == -1 && u.parent == v {
			VisitVertex2(vs, u, component)
		}
	}
}

func PrintGraph(vs Vertices, c *Component) {
	fmt.Println("graph {")
	for _, v := range vs {
		fmt.Printf("\t%d", v.x)
		if v.comp == c.num {
			fmt.Print(" [color=red]")
		}
		fmt.Println()
	}
	for _, v := range vs {
		for e := v.l.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			if u.x > v.x {
				fmt.Printf("\t%d--%d", v.x, u.x)
				if u.comp == v.comp && u.comp == c.num {
					fmt.Print(" [color=red]")
				}
				fmt.Println()
			}
		}
	}
	fmt.Println("}")
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

	bufstdin := bufio.NewReader(os.Stdin)
	for i := 0; i < M; i++ {
		var v, u int
		fmt.Fscan(bufstdin, &v, &u)
		vs[v].l.PushBack(vs[u])
		vs[u].l.PushBack(vs[v])
	}

	q := InitQueue(N)
	DFS1(vs, q)
	c := DFS2(vs, q)

	//PrintVertices(&vs)

	cs := make([]*Component, c)
	for i := 0; i < c; i++ {
		cs[i] = InitComponent(i)
	}

	for _, v := range vs {
		k := cs[v.comp]
		if k.vCount == 0 {
			k.repr = v
		}
		k.vCount++
		k.eCount += float32(v.l.Len()) / 2
	}

	//PrintComponents(&cs)

	maxC := cs[0]
	for i := 1; i < c; i++ {
		k := cs[i]
		if k.vCount > maxC.vCount {
			maxC = k
		} else if k.vCount == maxC.vCount {
			if k.eCount > maxC.eCount {
				maxC = k
			} else if k.eCount == maxC.eCount {
				if k.repr.x < maxC.repr.x {
					maxC = k
				}
			}
		}
	}

	PrintGraph(vs, maxC)
}
