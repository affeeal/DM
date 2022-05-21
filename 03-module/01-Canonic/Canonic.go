package main

import (
	"fmt"
	"sort"
)

var NUM int

const (
	WHITE = -1
	GRAY  = 0
	BLACK = 1
)

type Vertex struct {
	num        int
	mark       int
	output     []string
	transition Vertices
}

type Vertices []*Vertex

func (vs Vertices) Len() int           { return len(vs) }
func (vs Vertices) Swap(i, j int)      { vs[i], vs[j] = vs[j], vs[i] }
func (vs Vertices) Less(i, j int) bool { return vs[i].num < vs[j].num }

func InitVertex(i int) *Vertex {
	v := new(Vertex)
	v.num = i
	v.mark = WHITE
	v.transition = make(Vertices, 0)
	v.output = make([]string, 0)
	return v
}

func InitVertices(n int) *Vertices {
	vs := make(Vertices, n)
	for i := 0; i < n; i++ {
		vs[i] = InitVertex(i)
	}
	return &vs
}

func VisitVertex(v *Vertex) {
	if v.mark == GRAY || v.mark == BLACK {
		return
	}
	v.mark = GRAY
	v.num = NUM
	NUM++
	for _, u := range v.transition {
		VisitVertex(u)
	}
	v.mark = BLACK
}

func PrintAutomata(n, m int, vs *Vertices) {
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Printf("%d\n", (*vs)[i].transition[j].num)
			} else {
				fmt.Printf("%d ", (*vs)[i].transition[j].num)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Printf("%s\n", (*vs)[i].output[j])
			} else {
				fmt.Printf("%s ", (*vs)[i].output[j])
			}
		}
	}
}

func main() {
	var n, m, q0 int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &q0)
	vs := InitVertices(n)
	var q int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &q)
			(*vs)[i].transition = append((*vs)[i].transition, (*vs)[q])
		}
	}
	var out string
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%s", &out)
			(*vs)[i].output = append((*vs)[i].output, out)
		}
	}
	NUM = 0
	VisitVertex((*vs)[q0])
	sort.Sort(*vs)
	PrintAutomata(n, m, vs)
}
