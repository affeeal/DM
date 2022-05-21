package main

import "fmt"

type Vertex struct {
	number     int
	transition Vertices
	output     []byte
}

type Vertices []*Vertex

const K int = 2

var (
	N     int
	M     int
	WORDS []string
)

func InitVertex(i int) *Vertex {
	v := new(Vertex)
	v.number = i
	v.transition = make(Vertices, K)
	v.output = make([]byte, K)
	return v
}

func Detour(v *Vertex, i int, word []byte) {
	next, sym := v.transition[i], v.output[i]
	if v == next && sym == '-' {
		return
	}
	if sym != '-' {
		word = append(word, sym)
		if !WordInArray(string(word)) {
			WORDS = append(WORDS, string(word))
		}
		if len(word) == M {
			return
		}
	}
	for j := 0; j < K; j++ {
		Detour(next, j, word)
	}
}

func WordInArray(word string) bool {
	for _, w := range WORDS {
		if word == w {
			return true
		}
	}
	return false
}

func main() {
	fmt.Scanf("%d", &N)
	vs := make(Vertices, N)
	for i := 0; i < N; i++ {
		vs[i] = InitVertex(i)
	}
	var q int
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			fmt.Scanf("%d", &q)
			vs[i].transition[j] = vs[q]
		}
	}
	var sym string
	for i := 0; i < N; i++ {
		for j := 0; j < K; j++ {
			fmt.Scanf("%s", &sym)
			vs[i].output[j] = sym[0]
		}
	}
	var q0 int
	fmt.Scanf("%d", &q0)
	fmt.Scanf("%d", &M)
	for i := 0; i < K; i++ {
		Detour(vs[q0], i, []byte{})
	}
	for _, w := range WORDS {
		fmt.Println(w)
	}
}
