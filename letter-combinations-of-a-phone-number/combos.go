package main

import "fmt"


var letters = map[string][]string {
	"2": {"a", "b", "c"},
	"3": {"d", "e", "f"},
	"4": {"g", "h", "i"},
	"5": {"j", "k", "l"},
	"6": {"m", "n", "o"},
	"7": {"p", "q", "r", "s"},
	"8": {"t", "u", "v"},
	"9": {"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	var combinations []string



	for _, char := range digits {
		combinations = buildCombinations(combinations, string(char))

	}

	return combinations
}

func buildCombinations(existing []string, next string) []string {
	var combos []string
	options := letters[next]

	if len(existing) == 0 {
		combos = options
	}

	for _, prev := range existing {
		for _, option := range options {
			combo := prev + option
			combos = append(combos, combo)
		}
	}

	return combos
}

func main()  {
	result := letterCombinations("23")
	fmt.Println(result)
}
