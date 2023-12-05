package main

import (
	"fmt"
)

// --------------------------------------------------
// Runtime - 43 ms (Beats 86.32% of users with Go) //
// Memory - 8 MB (Beats 60.79% of users with Go)   //
// --------------------------------------------------
func firstMissingPositive02(nums []int) int {
	n := len(nums)
	fmt.Print("\n-----------------------------------------------------------")
	fmt.Printf("\nBase: %v\n", nums)

	for i := 0; i < n; i++ {
		fmt.Print("-----------------------------------------------------------\n")
		fmt.Printf("Processo: %v - index(%02d)valor(%02d)\n", nums, i, nums[i])

		for nums[i] > 0 && nums[i] <= n && nums[i] != nums[nums[i]-1] {
			fmt.Printf("Processo: %v - index(%02d)valor(%02d) ~> index(%02d)valor(%02d)\n", nums, i, nums[i], nums[i]-1, nums[nums[i]-1])
			nums[nums[i]-1], nums[i] = nums[i], nums[nums[i]-1]
		}
	}
	fmt.Print("-----------------------------------------------------------\n")
	fmt.Printf("Final: %v\n\n", nums)

	for i := 0; i < n; i++ {
		if nums[i] != i+1 {
			return i + 1
		}
	}

	return n + 1
}

// --------------------------------------------------
// Runtime - 39ms (Beats 94.47% of users with Go)  //
// Memory - 7.74MB (Beats 86.05% of users with Go) //
// --------------------------------------------------
func firstMissingPositive01(nums []int) int {
	fmt.Print("\n-----------------------------------------------------------")
	fmt.Printf("\nBase: %v\n", nums)

	for i := 0; i < len(nums); {
		fmt.Print("-----------------------------------------------------------\n")
		fmt.Printf("Processo: %v - index(%02d)valor(%02d)\n", nums, i, nums[i])

		cur := nums[i]
		if cur > 0 && cur <= len(nums) && nums[cur-1] != cur {
			fmt.Printf("Processo: %v - index(%02d)valor(%02d) ~> index(%02d)valor(%02d)\n", nums, i, nums[i], nums[i]-1, nums[nums[i]-1])
			nums[i], nums[nums[i]-1] = nums[nums[i]-1], nums[i]
			continue
		}

		i++
	}

	fmt.Print("-----------------------------------------------------------\n")
	fmt.Printf("Final: %v\n\n", nums)

	for i := 1; i < len(nums)+1; i++ {
		if nums[i-1] != i {
			return i
		}
	}

	return len(nums) + 1
}

// --------------------------------------------------
// Runtime - 38ms (Beats 95.26% of users with Go)  //
// Memory - 8.03MB (Beats 60.79% of users with Go) //
// --------------------------------------------------
func firstMissingPositive(nums []int) int {
	fmt.Print("\n-----------------------------------------------------------")
	fmt.Printf("\nBase: %v\n", nums)
	fmt.Print("-----------------------------------------------------------\n")

	index := make([]bool, len(nums)+1)
	for i, v := range nums {
		fmt.Printf("Processo: index(%02d)valor(%02d) %t\n", i, v, v > 0 && v <= len(nums))

		if v > 0 && v <= len(nums) {
			index[v] = true
		}
	}

	fmt.Print("\nindex slice:")
	for i, value := range index {
		fmt.Printf(" %02d %t", i, value)
	}

	fmt.Print("\n-----------------------------------------------------------\n")
	fmt.Print("Final: ")

	for i := 1; i <= len(nums); i++ {
		fmt.Printf("%02d(%v)", i, index[i])
		if !index[i] {
			fmt.Printf(" \n\n")
			return i
		}
		fmt.Print(" - ")
	}

	fmt.Printf(" \n\n")
	return len(nums) + 1
}

func process(nums []int, result int) {
	r := firstMissingPositive(nums)
	if r == result {
		fmt.Print("[OK]")
	} else {
		fmt.Print("[ERRO]")
	}

	fmt.Printf(" esperado(%02d)devolutiva(%02d)\n\n", result, r)
}

func main() {
	process(
		[]int{4, -1, 3, 2, 5, 1},
		6,
	)

	process(
		[]int{3, 4, -1, 1},
		2,
	)

	process(
		[]int{7, 8, 9, 11, 12},
		1,
	)

	process(
		[]int{1, 2, 7, 8, 9, -1, 11, 12, 3, 4, 5},
		6,
	)
}
