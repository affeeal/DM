package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
	"strconv"
)

// cycle cases need to be fixed

var (
	SRC []byte

	BEGIN int
	END   int
	POS   int
	SYM   byte

	PV *Vertex //  previous vertex
	CV *Vertex // current vertex
	VS []*Vertex

	ES []*Edge
)

// used structures

type Vertex struct {
	name     string
	dur      int
	max      int
	isInCP   bool
	parents  *list.List
	children *list.List
}

type Edge struct {
	v, u *Vertex
}

// Vertex functions

func InitVertex(name string, dur int) *Vertex {
	v := new(Vertex)
	v.name = name
	v.dur = dur
	v.max = -1
	v.isInCP = false
	v.parents = list.New()
	v.children = list.New()
	return v
}

func Find(name string) *Vertex {
	for _, v := range VS {
		if v.name == name {
			return v
		}
	}
	return nil
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

// Edge functions

func InitEdge(v, u *Vertex) *Edge {
	e := new(Edge)
	e.v = v
	e.u = u
	return e
}

func EdgeInES(v, u *Vertex) bool {
	for _, e := range ES {
		if e.v == v && e.u == u {
			return true
		}
	}
	return false
}

// auxiliary BNF functions

func IsSymLetter() bool {
	return SYM >= 'a' && SYM <= 'z' || SYM >= 'A' && SYM <= 'Z'
}

func IsSymNumber() bool {
	return SYM >= '0' && SYM <= '9'
}

func NextSym() {
	if POS+1 != END {
		POS++
		SYM = SRC[POS]
	} else if SYM != ';' {
		Error("expected ';'")
	}
}

// BNF
// <sentence> ::= <jobs> ';'
// <jobs> ::= <job> | <job> '<' <jobs>
// <job> ::= <name> ( <duration> ) | <name>

func Sentence() {
	Jobs()
	if SYM == ';' {
	} else {
		Error("expected ';'")
	}
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
		dur := Duration()
		CV = InitVertex(name, dur)
		VS = append(VS, CV)
		if SYM == ')' {
			NextSym()
		} else {
			Error("expected ')'")
		}
	}
	if !isInitialisation {
		CV = Find(name)
	}

	if PV != nil {
		CV.parents.PushBack(PV)
		PV.children.PushBack(CV)
	}
	PV = CV
}

func Name() string {
	if IsSymLetter() {
		var name []byte
		for IsSymLetter() || IsSymNumber() {
			name = append(name, SYM)
			NextSym()
		}
		return string(name)
	} else {
		Error("expected letter")
		return ""
	}
}

func Duration() int {
	if IsSymNumber() {
		var duration []byte
		for IsSymNumber() {
			duration = append(duration, SYM)
			NextSym()
		}
		dur, _ := strconv.Atoi(string(duration))
		return dur
	} else {
		Error("expected number")
		return -1
	}
}

// auxiliary functions

func Error(msg string) {
	fmt.Printf("error: %s\n", msg)
	os.Exit(1)
}

func PrintVertices() {
	for _, v := range VS {
		fmt.Printf("%s(%d): max == %d\n", v.name, v.dur, v.max)
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

func SetMaxPaths(p *Vertex) { // p is parent
	for e := p.children.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Vertex)

		var max *Vertex = nil
		VertexIsNotReady := false
		for e := v.parents.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)

			if max == nil || u.max > max.max {
				max = u
			} else if u.max == -1 {
				VertexIsNotReady = true
				break
			}
		}
		if !VertexIsNotReady {
			v.max = v.dur + max.max
			SetMaxPaths(v)
		}
	}
}

func BuildCMP(c *Vertex) { // c is child
	for e := c.parents.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Vertex)

		if v.max+c.dur == c.max {
			ES = append(ES, InitEdge(v, c))
			v.isInCP = true
			BuildCMP(v)
		}
	}
}

func PrintGraph() {
	fmt.Println("digraph {")
	for _, v := range VS {
		fmt.Printf("\t%s", v.name)
		if v.isInCP {
			fmt.Print(" [color=red]")
		}
		fmt.Println()
	}
	for _, v := range VS {
		for e := v.children.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			fmt.Printf("\t%s->%s", v.name, u.name)
			if EdgeInES(v, u) {
				fmt.Print(" [color=red]")
			}
			fmt.Println()
		}
	}
	fmt.Println("}")
}

func main() {
	BEGIN = 0
	var sym byte
	for {
		_, err := fmt.Scanf("%c", &sym)
		if err == io.EOF {
			break
		}
		if sym != ' ' && sym != '\n' {
			SRC = append(SRC, sym)
		}
		if sym == ';' {
			POS = BEGIN
			SYM = SRC[BEGIN]
			END = len(SRC)
			PV, CV = nil, nil
			Sentence()
			BEGIN = END
		}
	}
	r := FindRoot()
	r.max = r.dur
	SetMaxPaths(r)

	e := FindEnd()
	e.isInCP = true
	BuildCMP(e)

	PrintGraph()
}
