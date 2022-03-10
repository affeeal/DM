package main

import "fmt"

var (
	N int
	matrix [][]int
)

// Печать матрицы
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

// Модуль числа
func abs(a int) int {

	if a > 0 {
		return a
	}
	return -a
}

// НОД
func gcd(a, b int) int {

	if b == 0 {
		return a
	}
	return gcd(b, a % b)
}

// НОК
func lcm(a, b int) int {

	return abs(a * b / gcd(a, b))
}

// Преобразование столбца матрицы к столбцу диагональной матрицы
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

func main() {

	fmt.Scanf("%d", &N)
	matrix = make([][]int, N)
	for i := range matrix {

		matrix[i] = make([]int, N + 1)
		for j := range matrix[i] {

			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	// Приведение матрицы СЛАУ к диагональному виду.
	for j := range matrix {
		for i := j; i < N; i++ {

			if matrix[i][j] != 0 {
				diagonalizeColumn(i, j)
				break
			}
		}
	}	

	// Проверка СЛАУ на совместность.
	for i := range matrix {

		if matrix[i][i] == 0 {
			fmt.Printf("No solution\n")
			return
		}
	}

	// Упрощение итоговой СЛАУ и печать решений.
	for i := range matrix {

		gcd := gcd(matrix[i][i], matrix[i][N])
		matrix[i][i] /= gcd
		matrix[i][N] /= gcd
		fmt.Printf("%d/%d\n", matrix[i][N], matrix[i][i])
	}
}