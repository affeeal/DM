package main

import "fmt"

type Pair struct {
	s string        // key
	x int           // value
}

type PairSlice []Pair

func (array *PairSlice) Assign(s string, x int) {

	*array = append(*array, Pair {s, x})
}

func (array *PairSlice) Lookup(s string) (x int, exists bool) {

	for _, p := range *array {
		if p.s == s {
			return p.x, true
		}
	}
	return 0, false
}

type AssocArray interface {

	Assign(s string, x int)
	Lookup(s string) (x int, exists bool)
}

func main() {

	var array PairSlice = PairSlice {}
	var I AssocArray = &array

	var cmd string
	toContinue := true
	for toContinue {

		fmt.Scanf("%s", &cmd)
		if cmd == "assign" {

			var s string
			var x int
			fmt.Scanf("%s", &s)
			fmt.Scanf("%d", &x)
			I.Assign(s, x)
		} else if cmd == "lookup" {

			var s string
			fmt.Scanf("%s", &s)
			x, exists := I.Lookup(s)
			fmt.Printf("%d, %s\n", x, exists)
		} else if cmd == "exit" {

			toContinue = false
		}

		fmt.Printf("changes: %q\n", array)
	}
}