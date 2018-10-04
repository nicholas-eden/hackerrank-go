package main

import "fmt"

func main()  {
	result := numberToWords(10045678909)
	fmt.Println(result)
}

var conversion = map[int]string{
	0: "Zero",
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
	10: "Ten",
	11: "Eleven",
	12: "Twelve",
	13: "Thirteen",
	14: "Fourteen",
	15: "Fifteen",
	16: "Sixteen",
	17: "Seventeen",
	18: "Eighteen",
	19: "Nineteen",
	20: "Twenty",
	30: "Thirty",
	40: "Forty",
	50: "Fifty",
	60: "Sixty",
	70: "Seventy",
	80: "Eighty",
	90: "Ninety",
	100: "Hundred",
	1000: "Thousand",
	1000000: "Million",
	1000000000: "Billion",
	1000000000000: "Trillion",
}

var scales = []int{
	1000000000000,
	1000000000,
	1000000,
	1000,
	100,
}

func numberToWords(num int) string {
	if word, ok := conversion[num]; ok && num < 21 {
		return word
	}

	for _, scale := range scales {
		if scale > num {
			continue
		}

		times := num / scale
		remaining := num % scale
		name := conversion[scale]

		if remaining > 0 {
			return numberToWords(times) + " " + name + " " + numberToWords(remaining)
		}
		return numberToWords(times) + " " + name
	}

	tens := (num / 10) * 10
	remaining := num % tens
	name := conversion[tens]
	if remaining > 0 {
		return name + " " + numberToWords(remaining)
	}

	return name
}
