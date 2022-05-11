package main

import (
	"container/list"
	"fmt"
	"io"
	"strconv"
)

var (
	TIME  int
	COUNT int
	SYM   byte
	PV    *Vertex //  previous vertex
	CV    *Vertex // current vertex
	VS    Vertices
)

// used structures

type Vertex struct {
	name              string
	time, max, status int
	comp, T1, low     int
	parents, children *list.List
}

// status:
// -1 == the vertex was initialised;
// 0  == the vertex is not in CP;
// 1  == the vertex is in CP.

type Vertices []*Vertex

type Stack struct {
	data     Vertices
	cap, top int
}

// Vertex functions

func InitVertex(name string, time int) *Vertex {
	v := new(Vertex)
	v.name = name
	v.time = time
	v.max, v.status = -1, -1
	v.comp, v.low, v.T1 = 0, 0, 0
	v.parents, v.children = list.New(), list.New()
	return v
}

func FindByName(name string) *Vertex {
	for _, v := range VS {
		if v.name == name {
			return v
		}
	}
	return nil // this should never happen
}

func FindRoot() *Vertex {
	for _, v := range VS {
		if v.parents.Front() == nil {
			return v
		}
	}
	return nil // this should never happen
}

func FindEnd() *Vertex {
	for _, v := range VS {
		if v.children.Front() == nil {
			return v
		}
	}
	return nil // this should never happen
}

// Stack functions

func InitStack(N int) *Stack {
	s := new(Stack)
	s.data = make(Vertices, N)
	s.cap, s.top = 0, 0
	return s
}

func Push(s *Stack, v *Vertex) {
	s.data[s.top] = v
	s.top++
}

func Pop(s *Stack) *Vertex {
	s.top--
	return s.data[s.top]
}

// BNF
// <sentences> ::= <sentence> <sentences> |
// <sentence> ::= <jobs> ';'
// <jobs> ::= <job> | <job> '<' <jobs>
// <job> ::= <name> ( <duration> ) | <name>

func Sentences() {
	Sentence()
	// SYM is ';'
	if NextSym() {
		PV, CV = nil, nil
		Sentences()
	}
}

func Sentence() {
	Jobs()
	// SYM is ';'
}

func Jobs() {
	Job()
	if SYM == '<' {
		NextSym()
		Jobs()
	}
}

func Job() {
	name := Name()
	isInitialisation := false
	if SYM == '(' {
		isInitialisation = true
		NextSym()
		duration := Time()
		// SYM is ')'
		NextSym()
		CV = InitVertex(name, duration)
		VS = append(VS, CV)
	}
	if !isInitialisation {
		CV = FindByName(name)
	}
	if PV != nil {
		CV.parents.PushBack(PV)
		PV.children.PushBack(CV)
	}
	PV = CV
}

func Name() string {
	// SYM is a letter
	var name []byte
	for IsSymLetter() || IsSymNumber() {
		name = append(name, SYM)
		NextSym()
	}
	return string(name)
}

func Time() int {
	// SYM is a timeAsNumber
	var time []byte
	for IsSymNumber() {
		time = append(time, SYM)
		NextSym()
	}
	timeAsNumber, _ := strconv.Atoi(string(time))
	return timeAsNumber
}

// auxiliary BNF functions

func IsSymLetter() bool {
	return SYM >= 'a' && SYM <= 'z' || SYM >= 'A' && SYM <= 'Z'
}

func IsSymNumber() bool {
	return SYM >= '0' && SYM <= '9'
}

func NextSym() bool {
	_, err := fmt.Scanf("%c", &SYM)
	if err == io.EOF {
		return false
	}
	if SYM == ' ' || SYM == '\n' {
		return NextSym()
	}
	return true
}

// auxiliary functions

func PrintVertices() {
	for _, v := range VS {
		fmt.Printf("%s(%d):", v.name, v.time)
		fmt.Printf(" comp == %d; max == %d; status == %d\n", v.comp, v.max, v.status)

		fmt.Print("\tparents:")
		for e := v.parents.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			fmt.Printf(" %s;", u.name)
		}
		fmt.Println()

		fmt.Print("\tchildren:")
		for e := v.children.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			fmt.Printf(" %s;", u.name)
		}
		fmt.Println()
	}
}

// main

func CheckForCycles() {
	TIME, COUNT = 1, 1
	Tarjan()
	comps := make([]list.List, COUNT-1)
	for _, v := range VS {
		comps[v.comp-1].PushBack(v)
	}
	for _, l := range comps {
		if l.Len() != 1 {
			for e := l.Front(); e != nil; e = e.Next() {
				u := e.Value.(*Vertex)
				u.status = 0
			}
		}
	}
}

func Tarjan() {
	s := InitStack(len(VS))
	for _, v := range VS {
		if v.T1 == 0 {
			VisitVertexTarjan(v, s)
		}
	}
}

func VisitVertexTarjan(v *Vertex, s *Stack) {
	v.T1, v.low = TIME, TIME
	TIME++
	Push(s, v)
	for e := v.children.Front(); e != nil; e = e.Next() {
		u := e.Value.(*Vertex)
		if u.T1 == 0 {
			VisitVertexTarjan(u, s)
		}
		if u.comp == 0 && v.low > u.low {
			v.low = u.low
		}
	}
	if v.T1 == v.low {
		for {
			u := Pop(s)
			u.comp = COUNT
			if u == v {
				break
			}
		}
		COUNT++
	}
}

func SetMaxPaths(p *Vertex) { // p is a parent
	for e := p.children.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Vertex)

		if v.status == 0 {
			continue
		}

		var max *Vertex = nil
		VertexIsNotReady := false
		for e := v.parents.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)

			if u.status == 0 {
				continue
			}

			if max == nil || u.max > max.max {
				max = u
			} else if u.max == -1 {
				VertexIsNotReady = true
				break
			}
		}
		if !VertexIsNotReady {
			v.max = v.time + max.max
			SetMaxPaths(v)
		}
	}
}

func BuildCP(c *Vertex) { // c is child
	for e := c.parents.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Vertex)

		if v.status == 0 {
			continue
		}

		if v.max+c.time == c.max {
			v.status = 1
			BuildCP(v)
		}
	}
}

func PrintGraph() {
	fmt.Println("digraph {")
	for _, v := range VS {
		fmt.Printf("\t%s", v.name)
		if v.status == 1 {
			fmt.Print(" [color=red]")
		}
		fmt.Println()
	}
	for _, v := range VS {
		for e := v.children.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			fmt.Printf("\t%s->%s", v.name, u.name)
			if v.status == 1 && u.status == 1 && u.max == u.time+v.max {
				fmt.Print(" [color=red]")
			}
			fmt.Println()
		}
	}
	fmt.Println("}")
}

func main() {
	PV, CV = nil, nil
	NextSym()

	Sentences()
	CheckForCycles()

	r := FindRoot()
	r.max = r.time
	SetMaxPaths(r)

	e := FindEnd()
	e.status = 1
	BuildCP(e)

	PrintGraph()
}
