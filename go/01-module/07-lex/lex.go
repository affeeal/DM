package main

import "fmt"

type AssocArray interface {

	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

func compareStrings(a, b string) int {

	var (
		lenA int = len(a)
		lenB int = len(b)
	)
	if lenA < lenB {

		return -1
	} else if lenA > lenB {

		return 1
	} else {

		for i := 0; i < lenA; i++ {

			if a[i] > b[i] {

				return -1
			} else if a[i] < b[i] {

				return 1
			}
		}
		return 0
	}
}

func lex(sentence string, array AssocArray) []int {

	var (
		begin, end int = -1, -1
		inWord bool = false
		x, sLen int = 1, len(sentence)
		result []int = []int {}
	)

	addToAssocArray := func() {

		s := string(sentence[begin:end])
		fmt.Printf("s: %s\n", s)
		value, exists := array.Lookup(s)
		fmt.Printf("%s, %x\n", exists, value)
		if exists {

			result = append(result, value)
		} else {

			array.Assign(s, x)
			result = append(result, x)
			x++
		}
	}

	for pos, sym := range sentence {

		//fmt.Printf("sym: %c, pos: %d\n", sym, pos)
		if sym != ' ' && inWord == false {

			inWord = true
			begin = pos
		} else if sym == ' ' && inWord == true {

			inWord = false
			end = pos 
			addToAssocArray()
		}
	}
	if begin > end {

		end = sLen
		addToAssocArray()
	}
	return result
}

func main() {

	//var sentence string = "alpha x1 beta alpha x1 y"
	var sentence string = "a b c b a"

	var tree *AVLTreeNode = InitAVLTree()
	var arrayAVLTree AssocArray = tree
	var resultAVLTree []int = lex(sentence, arrayAVLTree)
	fmt.Printf("%d\n", resultAVLTree)

	/*l *SkipList = InitSkipList()
	arraySL AssocArray = l
	resultSL := lex(sentence, arraySL)
	fmt.Printf("%d\n", resultSL)*/
}