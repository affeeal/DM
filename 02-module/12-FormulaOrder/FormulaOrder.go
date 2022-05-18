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
	NUMBER Tag = 1 << iota // 1
	IDENT                  // 2
	PLUS                   // 4
	MINUS                  // 8
	MUL                    // 16
	DIV                    // 32
	LPAREN                 // 64
	RPAREN                 // 128
	COMMA                  // 256
	EQUAL                  // 512
)

var (
	FORMULAS     []byte
	SYM          byte
	FORMULAS_LEN int

	POS int

	TOKENS            []Token
	TOKEN             Token
	TOKENS_LEN        int
	NUM_OF_LEFT_ARGS  int
	NUM_OF_RIGHT_ARGS int
)

// scanning

func scan_formulas() {
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

func tokenize() {
	POS = -1
	next_sym()
	for in_formulas_bounds() {
		if sym_is_a_letter() {
			add_token(IDENT, ident())
		} else if sym_is_a_number() {
			add_token(NUMBER, number())
		} else if SYM == '+' {
			add_token(PLUS, string(SYM))
			next_sym()
		} else if SYM == '-' {
			add_token(MINUS, string(SYM))
			next_sym()
		} else if SYM == '*' {
			add_token(MUL, string(SYM))
			next_sym()
		} else if SYM == '/' {
			add_token(DIV, string(SYM))
			next_sym()
		} else if SYM == '(' {
			add_token(LPAREN, string(SYM))
			next_sym()
		} else if SYM == ')' {
			add_token(RPAREN, string(SYM))
			next_sym()
		} else if SYM == ',' {
			add_token(COMMA, string(SYM))
			next_sym()
		} else if SYM == '=' {
			add_token(EQUAL, string(SYM))
			next_sym()
		} else if SYM == '\n' {
			next_sym()
		} else {
			syntax_error()
		}
	}
	TOKENS_LEN = len(TOKENS)
}

func next_sym() {
	POS++
	if in_formulas_bounds() {
		SYM = FORMULAS[POS]
	}
}

func in_formulas_bounds() bool {
	return POS != FORMULAS_LEN
}

func add_token(tag Tag, image string) {
	TOKENS = append(TOKENS, Token{tag, image})
}

func sym_is_a_letter() bool {
	return SYM >= 'a' && SYM <= 'z' || SYM >= 'A' && SYM <= 'Z'
}

func sym_is_a_number() bool {
	return SYM >= '0' && SYM <= '9'
}

func ident() string {
	// SYM is a letter
	var new_ident []byte
	for in_formulas_bounds() && (sym_is_a_letter() || sym_is_a_number()) {
		new_ident = append(new_ident, SYM)
		next_sym()
	}
	return string(new_ident)
}

func number() string {
	// SYM is a number
	var new_number []byte
	for in_formulas_bounds() && sym_is_a_number() {
		new_number = append(new_number, SYM)
		next_sym()
	}
	return string(new_number)
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

func parse() {
	POS = -1
	next_token()
	formulas()
}

func formulas() {
	if in_tokens_bounds() {
		NUM_OF_LEFT_ARGS = 1
		NUM_OF_RIGHT_ARGS = 1
		formula()
		if NUM_OF_LEFT_ARGS != NUM_OF_RIGHT_ARGS {
			syntax_error()
		}
		formulas()
	}
}

func formula() {
	ident_list()
	if TOKEN.Tag&EQUAL == 0 {
		syntax_error()
	}
	next_token()
	expr_list()
}

func ident_list() {
	if TOKEN.Tag&IDENT == 0 {
		syntax_error()
	}
	next_token()
	another_ident()
}

func another_ident() {
	if TOKEN.Tag&COMMA != 0 {
		NUM_OF_LEFT_ARGS++
		next_token()
		ident_list()
	}
}

func expr_list() {
	expr()
	another_expr()
}

func another_expr() {
	if TOKEN.Tag&COMMA != 0 {
		NUM_OF_RIGHT_ARGS++
		next_token()
		expr_list()
	}
}

func expr() {
	term()
	expr_()
}

func expr_() {
	if TOKEN.Tag&PLUS != 0 {
		next_token()
		term()
		expr_()
	} else if TOKEN.Tag&MINUS != 0 {
		next_token()
		term()
		expr_()
	}
}

func term() {
	factor()
	term_()
}

func term_() {
	if TOKEN.Tag&MUL != 0 {
		next_token()
		factor()
		term_()
	} else if TOKEN.Tag&DIV != 0 {
		next_token()
		factor()
		term_()
	}
}

func factor() {
	if TOKEN.Tag&NUMBER != 0 {
		next_token()
	} else if TOKEN.Tag&IDENT != 0 {
		next_token()
	} else if TOKEN.Tag&LPAREN != 0 {
		next_token()
		expr()
		if TOKEN.Tag&RPAREN == 0 {
			syntax_error()
		}
		next_token()
	} else if TOKEN.Tag&MINUS != 0 {
		next_token()
		factor()
	} else {
		syntax_error()
	}
}

func next_token() {
	POS++
	if POS != TOKENS_LEN {
		TOKEN = TOKENS[POS]
	}
}

func in_tokens_bounds() bool {
	return POS != TOKENS_LEN
}

func syntax_error() {
	fmt.Println("syntax error")
	os.Exit(0)
}

// main

func main() {
	scan_formulas()
	tokenize()
	parse()
}
