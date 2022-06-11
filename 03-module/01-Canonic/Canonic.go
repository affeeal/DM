package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
)

var NUM int

const (
	WHITE = 0
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

func VisitVertex(v *Vertex) {
	if v.mark == BLACK {
		return
	}
	v.mark = BLACK
	v.num = NUM
	NUM++
	for _, u := range v.transition {
		VisitVertex(u)
	}
}

func main() {
	bufstdin := bufio.NewReader(os.Stdin)

	var n, m, q0 int
	fmt.Fscan(bufstdin, &n, &m, &q0)

	vs := make(Vertices, n)
	for i := 0; i < n; i++ {
		vs[i] = new(Vertex)
		vs[i].num = i
		vs[i].mark = WHITE
		vs[i].transition = make(Vertices, 0)
		vs[i].output = make([]string, 0)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var q int
			fmt.Fscan(bufstdin, &q)
			vs[i].transition = append(vs[i].transition, vs[q])
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			var out string
			fmt.Fscan(bufstdin, &out)
			vs[i].output = append(vs[i].output, out)
		}
	}

	NUM = 0
	VisitVertex(vs[q0])
	sort.Sort(vs)

	// printing automata
	fmt.Println(n)
	fmt.Println(m)
	fmt.Println(0)
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Printf("%d\n", vs[i].transition[j].num)
			} else {
				fmt.Printf("%d ", vs[i].transition[j].num)
			}
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			if j == m-1 {
				fmt.Printf("%s\n", vs[i].output[j])
			} else {
				fmt.Printf("%s ", vs[i].output[j])
			}
		}
	}
}
