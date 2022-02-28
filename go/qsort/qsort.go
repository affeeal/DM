package main

func partition(low, high int, 
	less func(i, j int) bool, 
	swap func(i, j int)) int {

	i := low
	j := low
	for j < high {
		if less(j, high) {
			swap(i, j)
			i++
		}
		j++
	}
	swap(i, high)
	return i
}

func qsortRec(low, high int, 
	less func(i, j int) bool, 
	swap func(i, j int)) {

	if low < high {
		q := partition(low, high, less, swap)
		qsortRec(low, q - 1, less, swap)
		qsortRec(q + 1, high, less, swap)
	}
}

func qsort(n int, 
	less func(i, j int) bool, 
	swap func(i, j int)) {

	qsortRec(0, n - 1, less, swap)
}

func main() {

}