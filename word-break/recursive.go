package main

import "fmt"

func main()  {
	result := wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"})
	//result := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	//result := wordBreak("applepenapple", []string{"apple", "pen"})

	fmt.Println(result)
}

func wordBreak(s string, wordDict []string) bool {
	byLength := make(map[int][]string)

	dictChars := make(map[rune]bool)

	for _, word := range wordDict {
		wl := len(word)
		byLength[wl] = append(byLength[wl], word)

		for _, char := range word {
			dictChars[char] = true
		}
	}

	for _, char := range s {
		if _, exists := dictChars[char]; exists == false {
			return false
		}
	}

	return search(s, &byLength)
}

func search(s string, byLength *map[int][]string) bool  {
	if len(s) == 0 {
		return true
	}

	for wl, words := range *byLength {
		if len(s) < wl {
			continue
		}

		candidate := s[0:wl]

		for _, word := range words {
			if candidate != word {
				continue
			}

			sub := s[wl:]
			result := search(sub, byLength)
			if result == false {
				continue
			}

			return true
		}

	}

	return false
}