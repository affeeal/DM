package main

import (
	"fmt"
	"math"
	"sort"
)

type Div struct {

	x, deg int
	divs []bool
}

type Divs []*Div

func (divs Divs) Len() int {

	return len(divs)
}

func (divs Divs) Swap(i, j int) {

	divs[i].x, divs[j].x = divs[j].x, divs[i].x
}

func (divs Divs) Less(i, j int) bool {

	return divs[i].x < divs[j].x
}

func InitDiv(i int) *Div {

	div := Div { i, 0, []bool {} }
	return &div
}

func InitGraph(n int) *[][]bool {

	graph := make([][]bool, n)
	for i := 0; i < n; i++ {

		graph[i] = make([]bool, i)
		for j := 0; j < i; j++ {

			graph[i][j] = false
		}
	}
	return &graph
}

func IsSubset(a, b *Div) bool {

	for i := range a.divs {

		if a.divs[i] == true && b.divs[i] != true {

			return false
		}
	}
	return true
}

func PrintGraph(divs *Divs, graph *[][]bool) {

	fmt.Println("graph {")
	for i := range *divs {

		fmt.Printf("\t%d\n", (*divs)[i].x)
	}
	for i := range *divs {

		for j := i + 1; j < (*divs).Len(); j++ {

			if (*graph)[j][i] == true {

				fmt.Printf("\t%d--%d\n", (*divs)[i].x, (*divs)[j].x)
			}
		}
	}
	fmt.Println("}")
}

func main() {

	var x int
	fmt.Scanf("%d", &x)
	divs := Divs {}
	start := int(math.Sqrt(float64(x)))

	for i := start ; i > 0 ; i-- {

		if x % i == 0 {

			divs = append(divs, InitDiv(i))
			if x != i * i {

				divs = append(divs, InitDiv(x / i))
			}
		}
	}

	sort.Sort(divs)
	len := divs.Len()
	for i := 0; i < len; i++ {

		for j := 0; j <= i; j++ {

			if divs[i].x % divs[j].x == 0 {

				divs[i].divs = append(divs[i].divs, true)
				divs[i].deg++
			} else {

				divs[i].divs = append(divs[i].divs, false)
			}
		}
	}

	graph := InitGraph(len)
	for i := 0; i < len; i++ {

		deg := 0
		for j := i + 1; j < len; j++ {

			if IsSubset(divs[i], divs[j]) {

				if deg == 0 {

					deg = divs[j].deg
					(*graph)[j][i] = true
				} else if divs[j].deg == deg {

					(*graph)[j][i] = true
				} else {

					break
				}
			}
		}
	}
	PrintGraph(&divs, graph)
}