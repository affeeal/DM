package main

import (
	"fmt"
	"os"
)

type Member struct {

	num int
	rel []*Member
}

func InitMember(N int) *Member {

	mem := Member { 0, make([]*Member, N)}
	for i := 0; i < N; i++ {

		mem.rel[i] = nil
	}
	return &mem
}

func BuildChain(mem *Member, N, t int) {

	mem.num = t
	for i := 0; i < N; i++ {

		cur := mem.rel[i]
		if cur != nil {
			if cur.num == 0 {
				BuildChain(cur, N, t ^ 3)
			} else if cur.num == t {
				fmt.Printf("No solution.\n")
				os.Exit(0)
			}
		}
	}
}

func NonPrevalentNum(nums []*int, N int) int {

	count := []int { 0, 0}
	for i := 0; i < N; i++ {

		if *nums[i] == 1 {
			count[0]++
		} else if *nums[i] == 2 {
			count[1]++
		}
	}
	if count[0] == count[1] {
		return 1
	} else {
		return 2
	}
}

func main() {

	var N int
	fmt.Scanf("%d", &N)
	mems := make([]*Member, N)
	nums := make([]*int, N)
	for i := 0; i < N; i++ {

		mems[i] = InitMember(N)
		nums[i] = &mems[i].num
	}
	var sym byte
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {

			fmt.Scanf("%c", &sym)
			if sym == '+' {
				mems[i].rel[j] = mems[j]
			}
		}
		fmt.Scanf("%c", &sym)
	}
	for i := 0; i < N; i++ {

		if mems[i].num == 0 {
			BuildChain(mems[i], N, NonPrevalentNum(nums, N))			
		}
	}
	t := NonPrevalentNum(nums, N) 
	for i := 0; i < N; i++ {

		if mems[i].num == t {
			fmt.Printf("%d ", i + 1)
		}
	}
	fmt.Println()
}