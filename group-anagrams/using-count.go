package main

import (
	"fmt"
)

func main() {
	result := groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"})
	fmt.Println(result)
}

func groupAnagrams(strs []string) [][]string {
	grouped := make(map[string][]string)

	for _, v := range strs {
		key := buildKey(v)

		grouped[key] = append(grouped[key], v)

	}

	var groupedSlice [][]string
	for _, group := range grouped {
		groupedSlice = append(groupedSlice, group)

	}

	return groupedSlice
}

var letters = []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
}

func buildKey(str string) string {
	chars := make(map[rune]int)
	for _, char := range str {
		chars[char] = chars[char] + 1
	}

	keySlice := make([]rune, len(str))
	for _, letter := range letters {
		count := chars[letter]

		for i := 0; i < count; i++ {
			keySlice = append(keySlice, letter)
		}

	}

	return string(keySlice)
}
