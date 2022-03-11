package main

import "fmt"

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
		/*fmt.Printf("\nsym: %c\n", sym)
		fmt.Printf("vStack: %q; vCount = %d\n", vStack, vCount)
		fmt.Printf("oStack: %q; oCount = %d\n", oStack, oCount)
		fmt.Printf("calculated: %q; cCount = %d\n", calculated, cCount)*/
	}
	return cCount
}

func main() {

	t1 := "x"
	t2 := "($xy)"
	t3 := "($(@ab)c)"
	t4 := "(#i($jk))"
	t5 := "(#($ab)($ab))"
	t6 := "(@(#ab)($ab))"
	t7 := "(#($a($b($cd)))(@($b($cd))($a($b($cd)))))"
	t8 := "(#($(#xy)($(#ab)(#ab)))(@z($(#ab)(#ab))))"
	
	fmt.Println(economPolish(t1))  //  =>  0
	fmt.Println(economPolish(t2))  //  =>  1
	fmt.Println(economPolish(t3))  //  =>  2
	fmt.Println(economPolish(t4))  //  =>  2
	fmt.Println(economPolish(t5))  //  =>  2
	fmt.Println(economPolish(t6))  //  =>  3
	fmt.Println(economPolish(t7))  //  =>  5
	fmt.Println(economPolish(t8))  //  =>  6
}