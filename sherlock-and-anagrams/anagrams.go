package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

var sortedSubs [][]string

func getSortedSub(s string, pos int, length int) string {
	zeroLength := length - 1
	sub := ""
	if len(sortedSubs) > pos {
		top := sortedSubs[pos]

		if len(top) > zeroLength {
			sub = top[zeroLength]
			if sub != "" {
				return sub
			}
		}
	}

	sub = s[pos : pos+length]

	chars := strings.Split(sub, "")
	sort.Strings(chars)
	sub = strings.Join(chars, "")
	btop := sortedSubs[pos]
	btop[zeroLength] = sub

	return sub
}

// Complete the sherlockAndAnagrams function below.
func sherlockAndAnagrams(s string) int32 {
	anagrams := int32(0)

	sortedSubs = make([][]string, len(s))

	for i := 0; i < len(s); i++ {
		sortedSubs[i] = make([]string, len(s)-i)
	}

	for i := 0; i < len(s); i++ {

		for j := i + 1; j < len(s); j++ {

			for size := 1; size <= len(s)-j; size++ {
				window := getSortedSub(s, i, size)
				compare := getSortedSub(s, j, size)

				if window == compare {
					anagrams++
				}

			}

		}
	}

	return anagrams
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := sherlockAndAnagrams(s)

		fmt.Fprintf(writer, "%d\n", result)
	}

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
