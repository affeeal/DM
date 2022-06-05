package main

import (
	"bufio"
	"fmt"
	"os"
)

func economPolish(expr string) int {
	var (
		oStack []string = make([]string, 128)
		vStack []string = make([]string, 128)
		calculated []string = []string {}
		vCount, oCount, cCount = 0, 0, 0
	)
	for pos := range expr {
		sym := expr[pos]
		if sym >= 'a' && sym <= 'z' {
			vStack[vCount] = string(sym)
			vCount++
		} else if sym == '#' || sym == '$' || sym == '@' {
			oStack[oCount] = string(sym)
			oCount++
		} else if sym == ')' {
			term := "(" + oStack[oCount - 1] + vStack[vCount - 2] + vStack[vCount - 1] + ")"
			isCalculated := false
			for i := range calculated {
				if calculated[i] == term {
					isCalculated = true
					break
				}
			}
			if !isCalculated {
				calculated = append(calculated, term)
				cCount++
			}
			vStack[vCount - 2], vStack[vCount - 1] = term, ""
			oStack[oCount - 1] = ""
			oCount--
			vCount--
		}
	}
	return cCount
}

func main() {
	expr, _ := bufio.NewReader(os.Stdin).ReadString('\n')
	fmt.Println(economPolish(expr))
}