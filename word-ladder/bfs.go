package main

import (
	"fmt"
)

func main() {
	result := ladderLength("qa", "sq", []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"})
	//result := ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log"})
	//result := ladderLength("a", "c", []string{"a", "b", "c",})
	fmt.Println(result)
}

type Node struct {
	key string
	depth int
	next *Node
}

func ladderLength(beginWord string, endWord string, wordList []string) int {
	contains := false
	for _, listed := range wordList {
		if listed == endWord {
			contains = true
			break
		}
	}

	if contains == false {
		return 0
	}

	visited := make(map[string]int)
	visited[beginWord] = 1
	head := &Node{
		beginWord,
		1,
		nil,
	}
	tail := head

	for head != nil {
		node := head
		head = node.next
		if head == nil {
			tail = nil
		}

		for _, word := range wordList {
			if !canTransform(node.key, word) {
				continue
			}

			if word == endWord {
				return node.depth + 1
			}

			depth, ok := visited[word]
			if depth > node.depth {
				visited[word] = node.depth + 1
			}

			if !ok {
				next := &Node{
					word,
					node.depth + 1,
					nil,
				}

				if tail != nil {
					tail.next = next
				}

				tail = next
				if head == nil {
					head = next
				}

			}

		}

	}


	return 0
}

func canTransform(word string, compare string) bool {
	diffCount := 0
	for i := range word {
		if word[i] != compare[i] {
			diffCount += 1

			if diffCount > 1 {
				return false
			}
		}

	}


	return true
}