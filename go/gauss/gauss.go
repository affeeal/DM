package main

import "fmt"

var (
	N int
	matrix [][]int
)

// Печать матрицы
func printMatrix() {

	fmt.Printf("\n")
	for i := 0; i < N; i++ {

		for j := 0; j < N + 1; j++ {

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

// Преобразование столба матрицы к столбу диагональной матрицы
func diagonalizeColumn(I, J int) {

	pivotElem := matrix[I][J]
	for i := 0; i < N; i++ {

		headElem := matrix[i][J]
		if headElem != 0 && i != I {

			lcm := lcm(headElem, pivotElem)
			//printMatrix()
			//fmt.Printf("lcm: %d\n", lcm)

			for j := 0; j < N + 1; j++ {

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
	for i := 0; i < N; i++ {

		matrix[i] = make([]int, N + 1)
		for j := 0; j < N + 1; j++ {

			fmt.Scanf("%d", &matrix[i][j])
		}
	}

	// Приведение матрицы СЛАУ к диагональному виду.
	for j := 0; j < N; j++ {

		for i := j; i < N; i++ {

			if matrix[i][j] != 0 {
				diagonalizeColumn(i, j)
				break
			}
		}
	}	

	// Проверка СЛАУ на совместность.
	for i := 0; i < N; i++ {

		if matrix[i][i] == 0 {

			fmt.Printf("No solution\n")
			return
		}
	}

	// Упрощение итоговой СЛАУ и печать решений.
	for i := 0; i < N; i++ {

		gcd := gcd(matrix[i][i], matrix[i][N])
		matrix[i][i] /= gcd
		matrix[i][N] /= gcd

		fmt.Printf("%d/%d\n", matrix[i][N], matrix[i][i])
	}
}