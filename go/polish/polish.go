package main

import "fmt"

func decodePolish(expr string) int {

	var (
		pos, len int = 0, len(expr)
		oStack []int = make([]int, 256)
		nStack []int = make([]int, 256)
		oCount, nCount int = 0, 0
	)

	for pos < len {

		sym := expr[pos]
		if sym >= '0' && sym <= '9' {

			nStack[nCount] = int(sym) - int('0')
			nCount++
		} else if sym == '*' {

			oStack[oCount] = int('*')
			oCount++
		} else if sym == '+' {

			oStack[oCount] = int('+')
			oCount++
		} else if sym == '-' {

			oStack[oCount] = int('-')
			oCount++
		} else if sym == ')' {

			if oStack[oCount - 1] == int('*') {

				nStack[nCount - 2] *= nStack[nCount - 1]
			} else if oStack[oCount - 1] == int('+') {

				nStack[nCount - 2] += nStack[nCount - 1]
			} else if oStack[oCount - 1] == int('-') {

				nStack[nCount - 2] -= nStack[nCount - 1]
			}
			oCount--
			nCount--
		}
		pos++
	}
	return nStack[0]
}

func main() {

	expr := "(* (* 4 (- (+ (* 2 3) (- 1 4)) 1)) (+ (+ (* 7 8) 4) (* 9 (- 2 8))))"
	fmt.Printf("%d\n", decodePolish(expr))  //  => 48
}