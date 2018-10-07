package main

import (
	"fmt"
	"sort"
	"strings"
)

func main()  {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(result)
}


func groupAnagrams(strs []string) [][]string {
	grouped := make(map[string][]string)

	for _, v := range strs {
		chars := strings.Split(v, "")
		sort.Strings(chars)

		key := strings.Join(chars, "")

		grouped[key] = append(grouped[key], v)

	}

	var groupedSlice [][]string
	for _, group := range grouped {
		groupedSlice = append(groupedSlice, group)

	}

	return groupedSlice
}