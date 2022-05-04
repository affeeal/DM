package main

import (
	"fmt"
	"os"
)

// global variables are declared with capital letters.

var (
	POS, LEN int
	SYM      byte
	STR      []byte
	TOKENS   []string
)

// auxiliary BNF functions

func is_letter() bool {
	return SYM >= 'a' && SYM <= 'z' || SYM >= 'A' && SYM <= 'Z'
}

func is_number() bool {
	return SYM >= '0' && SYM <= '9'
}

func is_comparison_op() (bool, string) {

	if SYM == '=' {
		return true, "="
	} else if SYM == '<' && has_next() {
		if is_next('>') {
			return true, "<>"
		} else if is_next('=') {
			return true, "<="
		}
		return true, "<"
	} else if SYM == '>' && has_next() {
		if is_next('=') {
			return true, ">="
		}
		return true, ">"
	}
	return false, ""
}

func has_next() bool {
	return POS+1 != LEN
}

func is_next(sym byte) bool {
	return STR[POS+1] == sym
}

func increase_pos(pos int) {
	for i := 0; i < pos; i++ {
		POS++
	}
	if POS != LEN {
		SYM = STR[POS]
	} else {
		if SYM != ';' {
			error()
		}
	}
}

func add_token(token string) {
	TOKENS = append(TOKENS, token)
}

func error() {
	error_debug()
	fmt.Println("error")
	os.Exit(1)
}

func error_debug() {
	fmt.Printf("POS: %d, SYM: %c\n", POS, SYM)
	fmt.Println(TOKENS)
}

// BNF

/*
<program> ::= <function> <program>

<function> ::= <ident> ( <formal-args-list> ) ':=' <expr> ;

<formal-args-list> ::= <ident-list> |
<ident-list> ::= <ident> | <ident> ',' <ident-list>

<expr> ::= <comparison_expr> '?' <comparison_expr> ':' <expr> | <comparison_expr>

<comparison_expr> ::= <arith_expr> <comparison_op> <arith_expr> | <arith_expr>
<comparison_op> ::= '=' | '<>' | '<' | '>' | '<=' | '>='

<arith_expr> ::= <term> <ARITH_EXPR>
<ARITH_EXPR> ::= '+' <term> <ARITH_EXPR> | '-' <term> <ARITH_EXPR> |

<term> ::= <factor> <TERM>
<TERM> ::= '*' <factor> <TERM> | '/' <factor> <TERM> |

<factor> ::=
  <number>
  | <ident>
  | <ident> ( <actual_args_list> )
  | ( <expr> )
  | '-' <factor>

<actual_args_list> ::= <expr-list> |
<expr-list> ::= <expr> | <expr> ',' <expr-list>
*/

func program() {
	function()
	if POS < LEN {
		program()
	}
}

func function() {

	ident()
	if SYM == '(' {
		add_token("(")
		increase_pos(1)

		formal_args_list()
		if SYM == ')' {
			add_token(")")
			increase_pos(1)

			if SYM == ':' && has_next() && is_next('=') {
				add_token(":=")
				increase_pos(2)

				expr()
				if SYM == ';' {
					add_token(";")
					increase_pos(1)
					return
				}
			}
		}
	}
	error()
}

func ident() {
	if is_letter() {
		var id []byte
		for is_letter() || is_number() {
			id = append(id, SYM)
			increase_pos(1)
		}
		add_token(string(id))
	} else {
		error()
	}
}

func formal_args_list() {

	if SYM == ')' {
		return
	}
	ident_list()
}

func ident_list() {
	ident()
	if SYM == ',' {
		add_token(",")
		increase_pos(1)

		ident_list()
	}
}

func expr() {
	comparison_expr()
	if SYM == '?' {
		add_token("?")
		increase_pos(1)

		comparison_expr()
		if SYM == ':' {
			add_token(":")
			increase_pos(1)

			expr()
			return
		}
		error()
	}
}

func comparison_expr() {
	arith_expr()
	verdict, op := is_comparison_op()
	if verdict {
		add_token(op)
		increase_pos(len(op))

		arith_expr()
	}
}

func arith_expr() {
	term()
	ARITH_EXPR()
}

func ARITH_EXPR() {
	if SYM == '+' || SYM == '-' {
		add_token(string(SYM))
		increase_pos(1)

		term()
		ARITH_EXPR()
	}
}

func term() {
	factor()
	TERM()
}

func TERM() {
	if SYM == '*' || SYM == '/' {
		add_token(string(SYM))
		increase_pos(1)

		factor()
		TERM()
	}
}

func factor() {
	if is_number() {
		number()
		return
	} else if SYM == '-' {
		add_token("-")
		increase_pos(1)

		factor()
		return
	} else if SYM == '(' {
		add_token("(")
		increase_pos(1)

		expr()

		if SYM == ')' {
			add_token(")")
			increase_pos(1)
			return
		}
		error()
	}
	ident()
	if SYM == '(' {
		add_token("(")
		increase_pos(1)

		actual_args_list()
		if SYM == ')' {
			add_token(")")
			increase_pos(1)
			return
		}
		error()
	}
}

func number() {
	var num []byte
	for is_number() {
		num = append(num, SYM)
		increase_pos(1)
	}
	add_token(string(num))
}

func actual_args_list() {
	if SYM == ')' {
		return
	}
	expr_list()
}

func expr_list() {
	expr()
	if SYM == ',' {
		add_token(",")
		increase_pos(1)

		expr_list()
	}
}

// main functions

func DeleteWhitespaces(src_str string) []byte {

	var temp_str []byte
	for _, sym := range []byte(src_str) {
		if sym != ' ' && sym != '\n' {
			temp_str = append(temp_str, sym)
		}
	}
	return temp_str
}

func main() {

	src_str := "fib(n) := fibrec(1,1,n);\nfibrec(a,b,n) := n=1 ? a : fibrec(b,a+b,n-1);"

	STR = DeleteWhitespaces(src_str)
	LEN = len(STR)
	POS = 0
	SYM = STR[0]

	program()

	fmt.Println(TOKENS)
}
