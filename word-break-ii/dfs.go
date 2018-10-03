package main

import (
	"fmt"
)

func main()  {
	//result := wordBreak("aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa", []string{"a","aa","aaa","aaaa","aaaaa","aaaaaa","aaaaaaa","aaaaaaaa","aaaaaaaaa","aaaaaaaaaa"})
	result := wordBreak("catsanddog", []string{"cat","cats","and","sand","dog"})
	//result := wordBreak("applepenapple", []string{"apple", "pen"})

	fmt.Println(result)
}

func wordBreak(s string, wordDict []string) []string {
	var matches []string
	byLength := make(map[int][]string)
	var queue *Node

	for _, word := range wordDict {
		wl := len(word)
		byLength[wl] = append(byLength[wl], word)
	}

	queue = &Node{
		0,
		"",
		nil,
	}

	for queue != nil {
		node := queue
		start := node.key
		queue = queue.next

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
					matches = append(matches, node.matched + word)
					continue
				}

				queue = &Node{
					start + wl,
					node.matched + word + " ",
					queue,
				}

			}

		}



	}

	return matches
}


type Node struct {
	key int
	matched string
	next *Node
}