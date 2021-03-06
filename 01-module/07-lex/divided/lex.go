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

		var s string = string(sentence[begin:end])
		value, exists := array.Lookup(s)
		if exists {

			result = append(result, value)
		} else {

			array.Assign(s, x)
			result = append(result, x)
			x++
		}
	}

	for pos, sym := range sentence {

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

	var (
		sentence string = "alpha x1 beta alpha x1 y "

		t *AVLTreeNode = InitAVLTree()
		arrayAVLTree AssocArray = t
		resultAVLTree []int = lex(sentence, arrayAVLTree)

		l *SkipList = InitSkipList()
		arraySL AssocArray = l
		resultSL = lex(sentence, arraySL)
	)

	fmt.Printf("SkipList: %d, AVLTree: %d.\n", resultSL, resultAVLTree)
}