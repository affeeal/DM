package main

import (
	"fmt"
	"math"
	"sort"
)

type Divs []int

func (divs Divs) Len() int           { return len(divs) }
func (divs Divs) Swap(i, j int)      { divs[i], divs[j] = divs[j], divs[i] }
func (divs Divs) Less(i, j int) bool { return divs[i] < divs[j] }

func InitGraph(n int) [][]bool {

	graph := make([][]bool, n)
	for i := 0; i < n; i++ {

		cap := i + 1
		graph[i] = make([]bool, cap)
		for j := 0; j < cap; j++ {

			graph[i][j] = false
		}
	}
	return graph
}

func PrintGraph(divs Divs, graph [][]bool) {

	fmt.Println("graph {")
	for _, d := range divs {

		fmt.Printf("\t%d\n", d)
	}
	for i, d := range divs {
		for j := i + 1; j < divs.Len(); j++ {

			if graph[j][i] == true {

				fmt.Printf("\t%d--%d\n", d, divs[j])
			}
		}
	}
	fmt.Println("}")
}

func main() {

	var x int
	fmt.Scanf("%d", &x)

	var divs Divs
	start := int(math.Sqrt(float64(x)))
	for i := start; i > 0; i-- {
		if x%i == 0 {
			divs = append(divs, i)
			if x != i*i {
				divs = append(divs, x/i)
			}
		}
	}

	sort.Sort(divs)
	len := divs.Len()

	graph := InitGraph(len)
	for i, divider := range divs {

		for j := i + 1; j < len; j++ {

			dividend := divs[j]
			if dividend%divider == 0 {

				isLeastDividend := true
				for k := i + 1; k < j; k++ {

					anotherDividend := divs[k]
					if dividend%anotherDividend == 0 && anotherDividend%divider == 0 {
						isLeastDividend = false
						break
					}
				}

				if isLeastDividend {
					graph[j][i] = true
				}
			}
		}
	}

	PrintGraph(divs, graph)
}
