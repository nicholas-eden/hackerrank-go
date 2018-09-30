package main

import "fmt"

func main()  {
	result := subsets([]int{9, 0, 3, 5, 7})
	fmt.Println(result)
}

func subsets(nums []int) [][]int {
	var sets [][]int
	sets = append(sets, []int{})

	for _, v := range nums {
		for _, set := range sets {
			added := append([]int{v}, set...)
			sets = append(sets, added)
		}
	}

	return sets
}
