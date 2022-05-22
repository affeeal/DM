package main

import (
	"fmt"
	"sort"
)

type Vertex struct {
	i          int
	mark       int
	parent     *Vertex
	depth      int
	transition Vertices
	output     []string
}

type Vertices []*Vertex

func (vs Vertices) Len() int           { return len(vs) }
func (vs Vertices) Swap(i, j int)      { vs[i], vs[j] = vs[j], vs[i] }
func (vs Vertices) Less(i, j int) bool { return vs[i].i < vs[j].i }

const (
	WHITE = 0
	BLACK = 1
)

var (
	N     int
	M     int
	COUNT int
)

func InitVertex(i int) *Vertex {
	v := new(Vertex)
	v.i = i
	v.mark = WHITE
	v.transition = make(Vertices, M)
	v.output = make([]string, M)
	return v
}

func SetCanonicalNumering(v *Vertex) {
	if v.mark == BLACK {
		return
	}
	v.mark = BLACK
	v.i = COUNT
	COUNT++
	for _, u := range v.transition {
		SetCanonicalNumering(u)
	}
}

func AufenkampHohn(Q Vertices) Vertices {
	m, pi := Split1(Q)
	for {
		var new_m int
		new_m, pi = Split(Q, pi)
		if m == new_m {
			break
		}
		m = new_m
	}
	new_Q := make(Vertices, 0)
	for _, q := range Q {
		new_q := pi[q.i]
		if !In(new_q, new_Q) {
			new_Q = append(new_Q, new_q)
			for k := 0; k < M; k++ {
				new_q.transition[k] = pi[q.transition[k].i]
				new_q.output[k] = q.output[k]
			}
		}
	}
	return new_Q
}

func Split1(Q Vertices) (int, Vertices) {
	m := N
	for _, q := range Q {
		q.parent = q
		q.depth = 0
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			q_1, q_2 := Q[i], Q[j]
			if Find(q_1) != Find(q_2) {
				eq := true
				for k := 0; k < M; k++ {
					if q_1.output[k] != q_2.output[k] {
						eq = false
						break
					}
				}
				if eq {
					Union(q_1, q_2)
					m--
				}
			}
		}
	}
	pi := make(Vertices, N)
	for _, q := range Q {
		pi[q.i] = Find(q)
	}
	return m, pi
}

func Find(x *Vertex) *Vertex {
	if x.parent == x {
		return x
	} else {
		x.parent = Find(x.parent)
		return x.parent
	}
}

func Union(x, y *Vertex) {
	root_x := Find(x)
	root_y := Find(y)
	if root_x.depth < root_y.depth {
		root_x.parent = root_y
	} else {
		root_y.parent = root_x
		if root_x.depth == root_y.depth && root_x != root_y {
			root_x.depth++
		}
	}
}

func Split(Q, pi Vertices) (int, Vertices) {
	m := N
	for _, q := range Q {
		q.parent = q
		q.depth = 0
	}
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			q_1, q_2 := Q[i], Q[j]
			if pi[q_1.i] == pi[q_2.i] && Find(q_1) != Find(q_2) {
				eq := true
				for k := 0; k < M; k++ {
					w_1, w_2 := q_1.transition[k].i, q_2.transition[k].i
					if pi[w_1] != pi[w_2] {
						eq = false
						break
					}
				}
				if eq {
					Union(q_1, q_2)
					m--
				}
			}
		}
	}
	for _, q := range Q {
		pi[q.i] = Find(q)
	}
	return m, pi
}

func In(new_q *Vertex, new_Q Vertices) bool {
	for _, q := range new_Q {
		if new_q == q {
			return true
		}
	}
	return false
}

func FindRoot(q_0 int, new_Q Vertices) *Vertex {
	for _, q := range new_Q {
		if q.i == q_0 {
			return q
		}
	}
	return nil // this should never happen
}

func PrintAutomata(Q Vertices) {
	fmt.Println("digraph {")
	fmt.Println("\trankdir = LR")
	for _, q := range Q {
		for j := 0; j < M; j++ {
			fmt.Printf("\t%d -> %d", q.i, q.transition[j].i)
			fmt.Printf(" [label = \"%c(%s)\"]\n", j+'a', q.output[j])
		}
	}
	fmt.Println("}")
}

func main() {
	var q_0 int
	fmt.Scanf("%d", &N)
	fmt.Scanf("%d", &M)
	fmt.Scanf("%d", &q_0)
	vs := make(Vertices, N)
	for i := 0; i < N; i++ {
		vs[i] = InitVertex(i)
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			var q int
			fmt.Scanf("%d", &q)
			vs[i].transition[j] = vs[q]
		}
	}
	for i := 0; i < N; i++ {
		for j := 0; j < M; j++ {
			var s string
			fmt.Scanf("%s", &s)
			vs[i].output[j] = s
		}
	}
	new_Q := AufenkampHohn(vs)
	COUNT = 0
	q := FindRoot(q_0, new_Q)
	SetCanonicalNumering(q)
	sort.Sort(new_Q)
	PrintAutomata(new_Q)
}
