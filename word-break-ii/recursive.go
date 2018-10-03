package main

import "fmt"

func main()  {
	//result := wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"})
	result := wordBreak("catsanddog", []string{"cat","cats","and","sand","dog"})
	//result := wordBreak("applepenapple", []string{"apple", "pen"})

	fmt.Println(result)
}


func wordBreak(s string, wordDict []string) []string {
	byLength := make(map[int][]string)

	for _, word := range wordDict {
		wl := len(word)
		byLength[wl] = append(byLength[wl], word)

	}

	cache := make(map[string][]string)
	return search(s, &byLength, &cache)
}

func search(s string, byLength *map[int][]string, cache *map[string][]string) []string  {
	var matches []string
	if len(s) == 0 {
		return matches
	}

	if m, found := (*cache)[s]; found {
		return m
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
			if len(sub) == 0 {
				matches = append(matches, candidate)
			}

			result := search(sub, byLength, cache)
			if len(result) == 0 {
				continue
			}

			for _, memo := range result {
				matches = append(matches, candidate + " " + memo)
			}

		}

	}

	(*cache)[s] = matches
	return matches
}