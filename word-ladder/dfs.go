package main

import (
	"fmt"
	"math"
)

func main() {
	result := ladderLength("qa", "sq", []string{"si","go","se","cm","so","ph","mt","db","mb","sb","kr","ln","tm","le","av","sm","ar","ci","ca","br","ti","ba","to","ra","fa","yo","ow","sn","ya","cr","po","fe","ho","ma","re","or","rn","au","ur","rh","sr","tc","lt","lo","as","fr","nb","yb","if","pb","ge","th","pm","rb","sh","co","ga","li","ha","hz","no","bi","di","hi","qa","pi","os","uh","wm","an","me","mo","na","la","st","er","sc","ne","mn","mi","am","ex","pt","io","be","fm","ta","tb","ni","mr","pa","he","lr","sq","ye"})
	//result := ladderLength("hit", "cog", []string{"hot","dot","dog","lot","log"})
	//result := ladderLength("1", "c", []string{"a", "b", "c",})
	fmt.Println(result)
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

	dfsCache := make(map[string]int)
	count := dfs(beginWord, endWord, wordList, &dfsCache, 1)
	return count
}

func dfs(beginWord string, endWord string, wordList []string, dfsCache *map[string]int, depth int) int {
	var successes []int
	if cached, found := (*dfsCache)[beginWord]; found {
		return cached
	}

	for i, candidate := range wordList {
		if candidate == beginWord {
			continue
		}
		if !canTransform(beginWord, candidate) {
			continue
		}

		if candidate == endWord {
			return depth + 1
		}

		listCopy := make([]string, len(wordList))
		copy(listCopy, wordList)
		subList := append(listCopy[:i], listCopy[i+1:]...)
		count := dfs(candidate, endWord, subList, dfsCache, depth + 1)
		if count > 0 {
			successes = append(successes, count)
		}
	}

	if len(successes) > 0 {
		lowest := math.MaxInt32
		for _, count := range successes {
			if count < lowest {
				lowest = count
			}
		}
		(*dfsCache)[beginWord] = lowest
		return lowest
	}

	return int(0)
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