package main

import "fmt"

type Vertex struct {
	number     int
	transition Vertices
	output     []string
}

type Vertices []*Vertex

func InitVertex(i, m int) *Vertex {
	v := new(Vertex)
	v.number = i
	v.transition = make(Vertices, m)
	v.output = make([]string, m)
	return v
}

func PrintGraph(n, m int, vs Vertices) {
	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("\t%d -> %d ", vs[i].number, vs[i].transition[j].number)
			fmt.Printf("[label = \"%c(%s)\"]\n", j+'a', vs[i].output[j])
		}
	}
	fmt.Println("}")
}

func main() {
	var n, m, q0 int
	fmt.Scanf("%d", &n)
	fmt.Scanf("%d", &m)
	fmt.Scanf("%d", &q0)
	vs := make(Vertices, n)
	for i := 0; i < n; i++ {
		vs[i] = InitVertex(i, m)
	}
	var q int
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%d", &q)
			vs[i].transition[j] = vs[q]
		}
	}
	var out string
	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Scanf("%s", &out)
			vs[i].output[j] = out
		}
	}
	PrintGraph(n, m, vs)
}
