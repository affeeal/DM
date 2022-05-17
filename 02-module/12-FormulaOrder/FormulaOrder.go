package main

import (
	"fmt"
	"io"
	"os"
)

type Tag int

type Token struct {
	Tag
	Image string
}

const (
	ERROR Tag = 1 << iota
	NUMBER
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

var (
	FORMULAS []byte
	LEN      int
	POS      int
	SYM      byte
	TOKENS   []Token
)

func scan() {
	var sym byte
	for {
		_, err := fmt.Scanf("%c", &sym)
		if err == io.EOF {
			return
		}
		if sym == ' ' {
			continue
		}
		FORMULAS = append(FORMULAS, sym)
	}
}

// indirect functions of lexical analysis

func add_token(tag Tag, image string) {
	TOKENS = append(TOKENS, Token{tag, image})
}

func next() {
	POS++
	if POS != LEN {
		SYM = FORMULAS[POS]
	}
}

func is_sym_a_letter() bool {
	return SYM >= 'a' && SYM <= 'z' || SYM >= 'A' && SYM <= 'Z'
}

func is_sym_a_number() bool {
	return SYM >= '0' && SYM <= '9'
}

func ident() {
	// SYM should be a letter
	if is_sym_a_letter() {
		var new_ident []byte
		for is_sym_a_letter() || is_sym_a_number() {
			new_ident = append(new_ident, SYM)
			next()
		}
		add_token(IDENT, string(new_ident))
	} else {
		syntax_error()
	}
}

func number() {
	// SYM is a number
	var new_number []byte
	for is_sym_a_number() {
		new_number = append(new_number, SYM)
		next()
	}
	add_token(NUMBER, string(new_number))
}

func syntax_error() {
	fmt.Println("syntax error")
	os.Exit(0)
}

// lexical analysis

func lexer() {
	POS = -1
	next()
	formulas()
}

func formulas() {
	if POS != LEN {
		formula()
		// SYM should be a '\n'
		next()
		formulas()
	}
}

func formula() {
	ident_list()
	if SYM != '=' {
		syntax_error()
	}
	add_token(EQUAL, "=")
	next()
	expr_list()
}

func ident_list() {
	ident()
	another_ident()
}

func another_ident() {
	if SYM == ',' {
		add_token(COMMA, ",")
		next()
		ident_list()
	}
}

func expr_list() {
	expr()
	another_expr()
}

func another_expr() {
	if SYM == ',' {
		add_token(COMMA, ",")
		next()
		expr_list()
	}
}

func expr() {
	term()
	expr_()
}

func expr_() {
	if SYM == '+' {
		add_token(PLUS, "+")
		next()
		term()
		expr_()
	} else if SYM == '-' {
		add_token(MINUS, "-")
		next()
		term()
		expr_()
	}
}

func term() {
	factor()
	term_()
}

func term_() {
	if SYM == '*' {
		add_token(MUL, "*")
		next()
		factor()
		term_()
	} else if SYM == '/' {
		add_token(DIV, "/")
		next()
		factor()
		term_()
	}
}

func factor() {
	if is_sym_a_number() {
		number()
	} else if is_sym_a_letter() {
		ident()
	} else if SYM == '(' {
		add_token(LPAREN, "(")
		next()
		expr()
		if SYM != ')' {
			syntax_error()
		}
		add_token(RPAREN, ")")
		next()
	} else if SYM == '-' {
		add_token(MINUS, "-")
		next()
		factor()
	}
}

// main

func main() {
	scan()
	LEN = len(FORMULAS)
	lexer()
}
