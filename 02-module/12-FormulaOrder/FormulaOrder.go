package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
)

// used structures

type Tag int

type Token struct {
	Tag
	image string
}

// Coordinates store information about the start and end positions
// of the formula corresponding to the vertex in the TOKENS array.

type Coordinates struct {
	start int
	end   int
}

type Colour int

type Vertex struct {
	image       string
	isDeclared  bool
	mark        Colour
	coordinates *Coordinates
	children    *list.List
}

type Vertices []*Vertex

// used constants and global variables

const (
	NUMBER Tag = 1 << iota
	IDENT
	PLUS
	MINUS
	MUL
	DIV
	LPAREN
	RPAREN
	COMMA
	EQUAL
)

const (
	WHITE Colour = -1 + iota
	GRAY
	BLACK
)

var (
	FORMULAS     []byte
	FORMULAS_LEN int
	SYM          byte

	POS int

	TOKENS     []Token
	TOKENS_LEN int
	TOKEN      Token

	GRAPH Vertices

	DEFINED_VERTICES Vertices
	DEFINED_VERTEX   *Vertex

	IDENT_LIST_LEN int
	EXPR_LIST_LEN  int
	COORDINATES    *Coordinates

	SORTED_LIST *list.List
)

// Vertex functions

func InitVertex(image string, isDeclared bool) *Vertex {
	v := new(Vertex)
	v.image = image
	v.mark = WHITE
	v.isDeclared = isDeclared
	v.coordinates = new(Coordinates)
	v.children = list.New()
	return v
}

func IsVertexInGraph(image string) (bool, *Vertex) {
	for _, v := range GRAPH {
		if v.image == image {
			return true, v
		}
	}
	return false, nil
}

func IsChild(v *Vertex) bool {
	for e := DEFINED_VERTEX.children.Front(); e != nil; e = e.Next() {
		u := e.Value.(*Vertex)
		if u.image == v.image {
			return true
		}
	}
	return false
}

// scanning

func ScanFormulas() {
	var sym byte
	for {
		_, err := fmt.Scanf("%c", &sym)
		if err == io.EOF {
			FORMULAS_LEN = len(FORMULAS)
			return
		}
		if sym != ' ' {
			FORMULAS = append(FORMULAS, sym)
		}
	}
}

// tokenizing

func Tokenize() {
	POS = -1
	NextSym()
	for InBoundsFORMUALS() {
		if SymIsALetter() {
			AddToken(IDENT, Ident())
		} else if SymIsANumber() {
			AddToken(NUMBER, Number())
		} else if SYM == '+' {
			AddToken(PLUS, string(SYM))
			NextSym()
		} else if SYM == '-' {
			AddToken(MINUS, string(SYM))
			NextSym()
		} else if SYM == '*' {
			AddToken(MUL, string(SYM))
			NextSym()
		} else if SYM == '/' {
			AddToken(DIV, string(SYM))
			NextSym()
		} else if SYM == '(' {
			AddToken(LPAREN, string(SYM))
			NextSym()
		} else if SYM == ')' {
			AddToken(RPAREN, string(SYM))
			NextSym()
		} else if SYM == ',' {
			AddToken(COMMA, string(SYM))
			NextSym()
		} else if SYM == '=' {
			AddToken(EQUAL, string(SYM))
			NextSym()
		} else if SYM == '\n' {
			NextSym()
		} else {
			SyntaxError()
		}
	}
	TOKENS_LEN = len(TOKENS)
}

func NextSym() {
	POS++
	if InBoundsFORMUALS() {
		SYM = FORMULAS[POS]
	}
}

func InBoundsFORMUALS() bool {
	return POS != FORMULAS_LEN
}

func AddToken(tag Tag, image string) {
	TOKENS = append(TOKENS, Token{tag, image})
}

func SymIsALetter() bool {
	return SYM >= 'a' && SYM <= 'z' || SYM >= 'A' && SYM <= 'Z'
}

func SymIsANumber() bool {
	return SYM >= '0' && SYM <= '9'
}

func Ident() string {
	// SYM is a letter
	var newIdent []byte
	for InBoundsFORMUALS() && (SymIsALetter() || SymIsANumber()) {
		newIdent = append(newIdent, SYM)
		NextSym()
	}
	return string(newIdent)
}

func Number() string {
	// SYM is a number
	var newNumber []byte
	for InBoundsFORMUALS() && SymIsANumber() {
		newNumber = append(newNumber, SYM)
		NextSym()
	}
	return string(newNumber)
}

// parsing

/*
<formulas> ::= <formula> <formulas> | Ɛ

<formula> ::= <ident_list> '=' <expr_list>

<ident_list> ::= <ident> <another_ident>
<another_ident> ::= ',' <ident_list> | Ɛ

<expr_list> ::= <expr> <another_expr>
<another_expr> ::= ',' <expr_list> | Ɛ

<expr> ::= <term> <expr_>
<expr_> ::= '+' <term> <expr_> | '-' <term> <expr_> | Ɛ
<term> ::= <factor> <term_>
<term_> ::= '*' <factor> <term_> | '/' <factor> <term_> | Ɛ
<factor> ::= <number> | <ident> | '(' <expr> ')' | '-' <factor>
*/

func Parse() {
	POS = -1
	NextToken()
	Formulas()
}

func Formulas() {
	if InBoundsTOKENS() {
		IDENT_LIST_LEN = 1
		EXPR_LIST_LEN = 1

		DEFINED_VERTICES = Vertices{}
		DEFINED_VERTEX = nil

		COORDINATES = new(Coordinates)
		COORDINATES.start = POS

		Formula()
		COORDINATES.end = POS - 1
		for _, v := range DEFINED_VERTICES {
			v.coordinates = COORDINATES
		}
		if IDENT_LIST_LEN != EXPR_LIST_LEN {
			SyntaxError()
		}
		Formulas()
	}
}

func Formula() {
	IdentList()
	if TOKEN.Tag&EQUAL == 0 {
		SyntaxError()
	}
	NextToken()
	ExprList()
}

func IdentList() {
	if TOKEN.Tag&IDENT == 0 {
		SyntaxError()
	}
	verdict, v := IsVertexInGraph(TOKEN.image)
	if verdict {
		if v.isDeclared {
			SyntaxError()
		} else {
			v.isDeclared = true
		}
	} else {
		v = InitVertex(TOKEN.image, true)
		GRAPH = append(GRAPH, v)
	}
	DEFINED_VERTICES = append(DEFINED_VERTICES, v)
	NextToken()
	AnotherIdent()
}

func AnotherIdent() {
	if TOKEN.Tag&COMMA != 0 {
		IDENT_LIST_LEN++
		NextToken()
		IdentList()
	}
}

func ExprList() {
	if EXPR_LIST_LEN > IDENT_LIST_LEN {
		SyntaxError()
	} else {
		DEFINED_VERTEX = DEFINED_VERTICES[EXPR_LIST_LEN-1]
	}
	Expr()
	AnotherExpr()
}

func AnotherExpr() {
	if TOKEN.Tag&COMMA != 0 {
		EXPR_LIST_LEN++
		NextToken()
		ExprList()
	}
}

func Expr() {
	Term()
	Expr_()
}

func Expr_() {
	if TOKEN.Tag&(PLUS|MINUS) != 0 {
		NextToken()
		Term()
		Expr_()
	}
}

func Term() {
	Factor()
	Term_()
}

func Term_() {
	if TOKEN.Tag&(MUL|DIV) != 0 {
		NextToken()
		Factor()
		Term_()
	}
}

func Factor() {
	if TOKEN.Tag&NUMBER != 0 {
		NextToken()
	} else if TOKEN.Tag&IDENT != 0 {
		verdict, v := IsVertexInGraph(TOKEN.image)
		if !verdict {
			v = InitVertex(TOKEN.image, false)
			GRAPH = append(GRAPH, v)
		}
		if !IsChild(v) {
			DEFINED_VERTEX.children.PushBack(v)
		}
		NextToken()
	} else if TOKEN.Tag&LPAREN != 0 {
		NextToken()
		Expr()
		if TOKEN.Tag&RPAREN == 0 {
			SyntaxError()
		}
		NextToken()
	} else if TOKEN.Tag&MINUS != 0 {
		NextToken()
		Factor()
	} else {
		SyntaxError()
	}
}

func NextToken() {
	POS++
	if POS != TOKENS_LEN {
		TOKEN = TOKENS[POS]
	}
}

func InBoundsTOKENS() bool {
	return POS != TOKENS_LEN
}

func SyntaxError() {
	fmt.Println("syntax error")
	os.Exit(0)
}

// checking Vertices

func CheckVerticesForDefinitions() {
	for _, v := range GRAPH {
		if !v.isDeclared {
			SyntaxError()
		}
	}
}

func DFS() {
	SORTED_LIST = list.New()
	for _, v := range GRAPH {
		VisitVertex(v)
	}
}

func VisitVertex(v *Vertex) {
	if v.mark == WHITE {
		v.mark = GRAY
		for e := v.children.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			VisitVertex(u)
		}
		v.mark = BLACK
		SORTED_LIST.PushFront(v)
	} else if v.mark == GRAY {
		Cycle()
	}
}

func Cycle() {
	fmt.Println("cycle")
	os.Exit(0)
}

// printing vertices

func PrintFormulas() {
	cs := new([]*Coordinates)
	for e := SORTED_LIST.Front(); e != nil; e = e.Next() {
		v := e.Value.(*Vertex)
		if FormulaIsNotPrinted(cs, v.coordinates) {
			PrintFormula(cs, v)
		}
	}
}

func FormulaIsNotPrinted(cs *[]*Coordinates, c *Coordinates) bool {
	for _, new_c := range *cs {
		if new_c == c {
			return false
		}
	}
	return true
}

func PrintFormula(cs *[]*Coordinates, v *Vertex) {
	for i := v.coordinates.start; i <= v.coordinates.end; i++ {
		if i != v.coordinates.end {
			fmt.Print(TOKENS[i].image + " ")
		} else {
			fmt.Println(TOKENS[i].image)
		}
	}
	*cs = append(*cs, v.coordinates)
}

// auxiliary functions

func PrintVertices() {
	for _, v := range GRAPH {
		fmt.Printf("%s(%t), {%d, %d}:", v.image, v.isDeclared, v.coordinates.start, v.coordinates.end)
		for e := v.children.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			fmt.Printf(" %s;", u.image)
		}
		fmt.Println()
	}
}

// main

func main() {
	ScanFormulas()
	Tokenize()
	Parse()
	CheckVerticesForDefinitions()
	DFS()
	PrintFormulas()
}
