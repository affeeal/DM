package main

import (
	"container/list"
	"fmt"
	"io"
	"os"
)

// global variables are declared with capital letters.

var (
	POS     int
	LEN     int
	SYM     byte
	PROGRAM []byte
	TOKENS  []string

	TIME  int
	COUNT int

	CV *Vertex // current vertex
	V  *Vertex

	VS Vertices
)

// used structures

type Stack struct {
	data     Vertices
	cap, top int
}

type Vertex struct {
	name          string
	comp, T1, low int
	children      *list.List
}

type Vertices []*Vertex

// Vertex functions

func InitVertex(name string) *Vertex {
	v := new(Vertex)
	v.name = name
	v.comp, v.low, v.T1 = 0, 0, 0
	v.children = list.New()
	return v
}

func Find() (bool, *Vertex) {
	for _, v := range VS {
		if v.name == TOKENS[len(TOKENS)-1] {
			return true, v
		}
	}
	return false, nil
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

// auxiliary BNF functions

func IsLetter() bool {
	return SYM >= 'a' && SYM <= 'z' || SYM >= 'A' && SYM <= 'Z'
}

func IsNumber() bool {
	return SYM >= '0' && SYM <= '9'
}

func IsComparisonOp() (bool, string) {
	if SYM == '=' {
		return true, "="
	} else if SYM == '<' && HasNext() {
		if IsNext('>') {
			return true, "<>"
		} else if IsNext('=') {
			return true, "<="
		}
		return true, "<"
	} else if SYM == '>' && HasNext() {
		if IsNext('=') {
			return true, ">="
		}
		return true, ">"
	}
	return false, ""
}

func HasNext() bool {
	return POS+1 != LEN
}

func IsNext(sym byte) bool {
	return PROGRAM[POS+1] == sym
}

func IncreasePos(pos int) {
	for i := 0; i < pos; i++ {
		POS++
	}
	if POS != LEN {
		SYM = PROGRAM[POS]
	} else {
		if SYM != ';' {
			Error()
		}
	}
}

func AddToken(token string) {
	TOKENS = append(TOKENS, token)
}

func Error() {
	fmt.Println("error")
	os.Exit(1)
}

// BNF

/*
<Program> ::= <Function> <Program>

<Function> ::= <Ident> ( <formal-args-list> ) ':=' <Expr> ;

<formal-args-list> ::= <Ident-list> |
<Ident-list> ::= <Ident> | <Ident> ',' <Ident-list>

<Expr> ::= <ComparisonExpr> '?' <ComparisonExpr> ':' <Expr> | <ComparisonExpr>

<ComparisonExpr> ::= <ArithExpr> <comparison_op> <ArithExpr> | <ArithExpr>
<comparison_op> ::= '=' | '<>' | '<' | '>' | '<=' | '>='

<ArithExpr> ::= <Term> <arithExpr>
<arithExpr> ::= '+' <Term> <arithExpr> | '-' <Term> <arithExpr> |

<Term> ::= <Factor> <term>
<term> ::= '*' <Factor> <term> | '/' <Factor> <term> |

<Factor> ::=
  <Number>
  | <Ident>
  | <Ident> ( <ActualArgsList> )
  | ( <Expr> )
  | '-' <Factor>

<ActualArgsList> ::= <Expr-list> |
<Expr-list> ::= <Expr> | <Expr> ',' <Expr-list>
*/

func Program() {
	Function()
	if POS < LEN {
		Program()
	}
}

func Function() {
	Ident()
	var VertexInVS bool
	VertexInVS, CV = Find()
	if !VertexInVS {
		CV = InitVertex(TOKENS[len(TOKENS)-1])
		VS = append(VS, CV)
	}
	if SYM == '(' {
		AddToken("(")
		IncreasePos(1)

		FormalArgsList()
		if SYM == ')' {
			AddToken(")")
			IncreasePos(1)

			if SYM == ':' && HasNext() && IsNext('=') {
				AddToken(":=")
				IncreasePos(2)

				Expr()
				if SYM == ';' {
					AddToken(";")
					IncreasePos(1)
					return
				}
			}
		}
	}
	Error()
}

func Ident() {
	if IsLetter() {
		var id []byte
		for IsLetter() || IsNumber() {
			id = append(id, SYM)
			IncreasePos(1)
		}
		AddToken(string(id))
	} else {
		Error()
	}
}

func FormalArgsList() {
	if SYM == ')' {
		return
	}
	IdentList()
}

func IdentList() {
	Ident()
	if SYM == ',' {
		AddToken(",")
		IncreasePos(1)

		IdentList()
	}
}

func Expr() {
	ComparisonExpr()
	if SYM == '?' {
		AddToken("?")
		IncreasePos(1)

		ComparisonExpr()
		if SYM == ':' {
			AddToken(":")
			IncreasePos(1)

			Expr()
			return
		}
		Error()
	}
}

func ComparisonExpr() {
	ArithExpr()
	verdict, op := IsComparisonOp()
	if verdict {
		AddToken(op)
		IncreasePos(len(op))

		ArithExpr()
	}
}

func ArithExpr() {
	Term()
	arithExpr()
}

func arithExpr() {
	if SYM == '+' || SYM == '-' {
		AddToken(string(SYM))
		IncreasePos(1)

		Term()
		arithExpr()
	}
}

func Term() {
	Factor()
	term()
}

func term() {
	if SYM == '*' || SYM == '/' {
		AddToken(string(SYM))
		IncreasePos(1)

		Factor()
		term()
	}
}

func Factor() {
	if IsNumber() {
		Number()
		return
	} else if SYM == '-' {
		AddToken("-")
		IncreasePos(1)

		Factor()
		return
	} else if SYM == '(' {
		AddToken("(")
		IncreasePos(1)

		Expr()

		if SYM == ')' {
			AddToken(")")
			IncreasePos(1)
			return
		}
		Error()
	}
	Ident()
	if SYM == '(' {
		var VertexInVS bool
		VertexInVS, V = Find()
		if !VertexInVS {
			V = InitVertex(TOKENS[len(TOKENS)-1])
			VS = append(VS, V)
		}
		CV.children.PushBack(V)

		AddToken("(")
		IncreasePos(1)

		ActualArgsList()
		if SYM == ')' {
			AddToken(")")
			IncreasePos(1)
			return
		}
		Error()
	}
}

func Number() {
	var num []byte
	for IsNumber() {
		num = append(num, SYM)
		IncreasePos(1)
	}
	AddToken(string(num))
}

func ActualArgsList() {
	if SYM == ')' {
		return
	}
	ExprList()
}

func ExprList() {
	Expr()
	if SYM == ',' {
		AddToken(",")
		IncreasePos(1)

		ExprList()
	}
}

// auxiliary functions

func PrintVertices() {
	for _, v := range VS {
		fmt.Printf("%s:", v.name)
		for e := v.children.Front(); e != nil; e = e.Next() {
			u := e.Value.(*Vertex)
			fmt.Printf(" %s;", u.name)
		}
		fmt.Println()
	}
}

// main functions

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

func main() {

	var sym byte
	for {
		_, err := fmt.Scanf("%c", &sym)
		if err == io.EOF {
			break
		}

		if sym != ' ' && sym != '\n' && sym != '\t' {
			PROGRAM = append(PROGRAM, sym)
		}
	}

	LEN = len(PROGRAM)
	POS = 0
	SYM = PROGRAM[POS]

	Program()

	TIME, COUNT = 1, 1
	Tarjan()
	PrintVertices()
	fmt.Println(COUNT - 1)
}
