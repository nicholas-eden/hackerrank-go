package main

import "fmt"

func lengthOfLongestSubstring(s string) int {
	counts := make(map[rune]int)
	longest := 0

	start := 0
	for i, char := range s {
		count, ok := counts[char]
		if ok {
			counts = make(map[rune]int)
			start = count + 1
			for j, charSub := range s[start:i] {
				counts[charSub] = j + start
			}

		}

		counts[char] = i

		if i - start + 1 > longest {
			longest = i+1 - start
		}
	}

	return longest
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abdvdf"))
}
