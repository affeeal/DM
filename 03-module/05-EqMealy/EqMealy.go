package main

import (
	"fmt"
	"os"
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
	COUNT int
)

func ScanAutomata(N, M, q0 *int, alphabet *[]string) Vertices {
	fmt.Scanf("%d", &*N)
	fmt.Scanf("%d", &*M)
	fmt.Scanf("%d", &*q0)
	vs := make(Vertices, *N)
	for i := 0; i < *N; i++ {
		vs[i] = InitVertex(i, *M)
	}
	for i := 0; i < *N; i++ {
		for j := 0; j < *M; j++ {
			var q int
			fmt.Scanf("%d", &q)
			vs[i].transition[j] = vs[q]
		}
	}
	for i := 0; i < *N; i++ {
		for j := 0; j < *M; j++ {
			var s string
			fmt.Scanf("%s", &s)
			if !InAlphabet(s, *alphabet) {
				*alphabet = append(*alphabet, s)
			}
			vs[i].output[j] = s
		}
	}
	return vs
}

func InitVertex(i int, M int) *Vertex {
	v := new(Vertex)
	v.i = i
	v.mark = WHITE
	v.transition = make(Vertices, M)
	v.output = make([]string, M)
	return v
}

func InAlphabet(new_s string, alphabet []string) bool {
	for _, s := range alphabet {
		if new_s == s {
			return true
		}
	}
	return false
}

func EqualAlphabets(alphabet_1, alphabet_2 []string) bool {
	if len(alphabet_1) != len(alphabet_2) {
		return false
	}
	for _, w1 := range alphabet_1 {
		isInAnotherAlphabet := false
		for _, w2 := range alphabet_2 {
			if w1 == w2 {
				isInAnotherAlphabet = true
				break
			}
		}
		if !isInAnotherAlphabet {
			return false
		}
	}
	return true
}

func AufenkampHohn(Q Vertices, N, M int) Vertices {
	m, pi := Split1(Q, N, M)
	for {
		var new_m int
		new_m, pi = Split(Q, pi, N, M)
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

func Split1(Q Vertices, N, M int) (int, Vertices) {
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

func Split(Q, pi Vertices, N, M int) (int, Vertices) {
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

func CanonizeAutomata(new_Q *Vertices, q0 int) {
	COUNT = 0
	q := FindRoot(q0, *new_Q)
	SetCanonicalNumering(q)
	sort.Sort(new_Q)
}

func FindRoot(q_0 int, new_Q Vertices) *Vertex {
	for _, q := range new_Q {
		if q.i == q_0 {
			return q
		}
	}
	return nil // this should never happen
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

func EqualAutomatas(Q1, Q2 Vertices, M int) bool {
	new_N1, new_N2 := len(Q1), len(Q2)
	if new_N1 != new_N2 {
		NotEqual()
	}
	for i := 0; i < new_N1; i++ {
		q1, q2 := Q1[i], Q2[i]
		for j := 0; j < M; j++ {
			if q1.transition[j].i != q2.transition[j].i ||
				q1.output[j] != q2.output[j] {
				return false
			}
		}
	}
	return true
}

func Equal() {
	fmt.Println("EQUAL")
}

func NotEqual() {
	fmt.Println("NOT EQUAL")
	os.Exit(0)
}

func main() {
	var N1, M1, q0_1 int
	var alphabet_1 []string
	Q1 := ScanAutomata(&N1, &M1, &q0_1, &alphabet_1)
	var N2, M2, q0_2 int
	var alphabet_2 []string
	Q2 := ScanAutomata(&N2, &M2, &q0_2, &alphabet_2)
	if !EqualAlphabets(alphabet_1, alphabet_2) {
		NotEqual()
	}
	new_Q1 := AufenkampHohn(Q1, N1, M1)
	CanonizeAutomata(&new_Q1, q0_1)
	new_Q2 := AufenkampHohn(Q2, N2, M2)
	CanonizeAutomata(&new_Q2, q0_2)
	if EqualAutomatas(new_Q1, new_Q2, M1) {
		Equal()
	} else {
		NotEqual()
	}
}
