package main

import (
	"fmt"
)

func main()  {
	//result := wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaab", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"})
	//result := wordBreak("catsandog", []string{"cats", "dog", "sand", "and", "cat"})
	result := wordBreak("applepenapple", []string{"apple", "pen"})

	fmt.Println(result)
}

func wordBreak(s string, wordDict []string) bool {
	byLength := make(map[int][]string)
	visited := make(map[int]bool)
	var queue *Node

	for _, word := range wordDict {
		wl := len(word)
		byLength[wl] = append(byLength[wl], word)
	}

	queue = &Node{
		0,
		nil,
	}

	for queue != nil {
		start := queue.value
		queue = queue.next

		if visited[start] == true {
			continue
		}

		visited[start] = true

		for wl, words := range byLength {
			if wl + start > len(s) {
				continue
			}

			candidate := s[start: start + wl]

			for _, word := range words {
				if candidate != word {
					continue
				}

				if start + wl == len(s) {
					return true
				}

				queue = &Node{
					start + wl,
					queue,
				}

			}

		}



	}

	return false
}


type Node struct {
	value int
	next *Node
}