package main

import "fmt"

var (
	N int
	matrix [][]int
)

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func gcd(a, b int) int {

	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

func lcm(a, b int) int {

	return abs(a * b / gcd(a, b))
}

func diagonalizeColumn(I, J int) {
	pivotElem := matrix[I][J]
	for i := range matrix {
		headElem := matrix[i][J]
		if headElem != 0 && i != I {
			lcm := lcm(headElem, pivotElem)
			for j := range matrix[0] {

				matrix[i][j] *= lcm / headElem
				matrix[i][j] -= lcm / pivotElem * matrix[I][j]
			}
		}
	}
	matrix[I], matrix[J] = matrix[J], matrix[I]
}

func printMatrix() {
	fmt.Printf("\n")
	for i := range matrix {
		for j := range matrix[0] {
			if j == N {
				fmt.Printf("%d\n", matrix[i][j])
			} else {
				fmt.Printf("%d ", matrix[i][j])
			}
		}
	}
}

func main() {
	fmt.Scanf("%d", &N)
	matrix = make([][]int, N)
	for i := range matrix {
		matrix[i] = make([]int, N + 1)
		for j := range matrix[i] {
			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	for j := range matrix {
		for i := j; i < N; i++ {
			if matrix[i][j] != 0 {
				diagonalizeColumn(i, j)
				break
			}
		}
	}	

	for i := range matrix {
		if matrix[i][i] == 0 {
			fmt.Printf("No solution\n")
			return
		}
	}

	for i := range matrix {
		gcd := gcd(matrix[i][i], matrix[i][N])
		matrix[i][i] /= gcd
		matrix[i][N] /= gcd
		if matrix[i][N] > 0 && matrix[i][i] < 0 {
			fmt.Printf("%d/%d\n", -matrix[i][N], -matrix[i][i])
		} else {
			fmt.Printf("%d/%d\n", matrix[i][N], matrix[i][i])
		}
	}
}