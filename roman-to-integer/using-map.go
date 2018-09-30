package main

import "fmt"

func main() {
	result := romanToInt("MCMXCIV")
	fmt.Println(result)
}

var convert = map[rune]int{
	'I': 1,
	'V': 5,
	'X': 10,
	'L': 50,
	'C': 100,
	'D': 500,
	'M': 1000,
}

func romanToInt(s string) int {
	total := 0

	var prev int
	for _, letter := range s {
		value := convert[letter]
		total += value
		if prev < value {
			total -= prev * 2
		}

		prev = value
	}

	return total
}
