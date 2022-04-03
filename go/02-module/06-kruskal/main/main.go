package main

import (
	"fmt"
	"math"
	"sort"
)

type Vertex struct {
	x, y   float64
	depth  int
	parent *Vertex
}

type Edge struct {
	u, v   *Vertex
	weight float64
}

type Edges []*Edge

func (es Edges) Len() int           { return len(es) }
func (es Edges) Swap(i, j int)      { es[i], es[j] = es[j], es[i] }
func (es Edges) Less(i, j int) bool { return es[i].weight < es[j].weight }

func InitVertex(x, y float64) *Vertex {

	u := new(Vertex)
	u.x = x
	u.y = y
	u.depth = 0
	u.parent = u
	return u
}

func InitEdge(u, v *Vertex) *Edge {

	e := new(Edge)
	e.u = u
	e.v = v
	e.weight = math.Sqrt(math.Pow(u.x-v.x, 2) + math.Pow(u.y-v.y, 2))
	return e
}

func Find(u *Vertex) *Vertex {

	if u.parent == u {

		return u
	}
	return Find(u.parent)
}

func Union(u, v *Vertex) {

	uRoot := Find(u)
	vRoot := Find(v)
	if uRoot.depth < vRoot.depth {

		uRoot.parent = vRoot
	} else {

		vRoot.parent = uRoot
		if uRoot.depth == vRoot.depth && uRoot != vRoot {

			uRoot.depth++
		}
	}
}

func SpanningTree(es Edges) {

	esResult := make(Edges, 0)
	for _, e := range es {

		if Find(e.u) != Find(e.v) {

			Union(e.u, e.v)
			esResult = append(esResult, e)
		}
	}
	var sum float64 = 0
	for _, e := range esResult {

		//PrintEdge(e)
		sum += e.weight
	}
	fmt.Printf("%.2f\n", sum)
}

/*func PrintEdge(e *Edge) {

	fmt.Printf("u: %.1f, %.1f, v: %.1f, %.1f; len: %.2f\n", e.u.x, e.u.y, e.v.x, e.v.y, e.weight)
}

func PrintEdges(es Edges) {

	for _, e := range es {
		PrintEdge(e)
	}
}*/

func main() {

	var N int
	fmt.Scanf("%d\n", &N)

	var x, y float64
	vs := make([]*Vertex, N)
	for i := 0; i < N; i++ {

		fmt.Scanf("%f %f\n", &x, &y)
		vs[i] = InitVertex(x, y)
	}

	es := make(Edges, 0)
	for i := 0; i < N-1; i++ {

		for j := i + 1; j < N; j++ {

			es = append(es, InitEdge(vs[i], vs[j]))
		}
	}

	sort.Sort(es)
	SpanningTree(es)
}
