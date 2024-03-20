package main

import "fmt"

func main() {
	nums := read()
	answer := solution(nums)
	fmt.Println(answer)
}

func read() []int {
	var a int
	nums := make([]int, 0, 10)
	for {
		fmt.Scan(&a)
		if a == -2*1000000000 {
			return nums
		}
		nums = append(nums, a)
	}
}

type seqType string

const (
	ConstType    seqType = "CONSTANT"
	AscType      seqType = "ASCENDING"
	WeakAscType  seqType = "WEAKLY ASCENDING"
	DescType     seqType = "DESCENDING"
	WeakDescType seqType = "WEAKLY DESCENDING"
	RandType     seqType = "RANDOM"
)

const (
	sameFlag = 1 << iota
	ascFlag
	descFlag
)

func solution(nums []int) seqType {
	if len(nums) == 0 {
		return ConstType
	}
	prev := nums[0]
	flags := 0
	for i := 1; i < len(nums); i++ {
		n := nums[i]
		switch {
		case n == prev:
			flags |= sameFlag
		case n > prev:
			flags |= ascFlag
		case n < prev:
			flags |= descFlag
		}
		prev = n
	}
	switch flags {
	case sameFlag:
		return ConstType
	case ascFlag:
		return AscType
	case ascFlag | sameFlag:
		return WeakAscType
	case descFlag:
		return DescType
	case descFlag | sameFlag:
		return WeakDescType
	default:
		return RandType
	}
}
