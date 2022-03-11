package main

import "fmt"

func sum(a, b []int32, p int) []int32 {

	var (
		sum []int32 = make([]int32, len(a))
		add int32 = 0
	)
	for i := range a {

		sum[i] = (a[i] + b[i] + add) % int32(p)
		add = (a[i] + b[i] + add) / int32(p)
	}
	if add != 0 {

		sum = append(sum, add)
	}
	return sum
}

func add(a, b []int32, p int) []int32 {

	var lenA, lenB int = len(a), len(b)
	if lenA < lenB {

		augA := make([]int32, lenB)
		for i, x := range a {

			augA[i] = x
		}
		return sum(augA, b, p)
	} else if lenA > lenB {

		augB := make([]int32, lenA)
		for i, x := range b {

			augB[i] = x
		}
		return sum(a, augB, p)
	} else {

		return sum(a, b, p)
	}
}

func main() {
	
	var (
		a, b []int32 = []int32 { 7, 1 }, []int32 { 2, 3, 1 } //  17 + 132 в 8-ной
		p int = 8
	)
	fmt.Println("%q", add(a, b, p))  //  => 151
}