package main

import "fmt"

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)

	for i, v := range nums {
		diff := target - v

		match, ok := m[diff]
		if ok {
			return []int{match, i}
		}

		m[v] = i
	}

	return []int{}
}

func main()  {
	result := twoSum([]int{2, 7, 11, 15}, 9)
	fmt.Println(result)
}