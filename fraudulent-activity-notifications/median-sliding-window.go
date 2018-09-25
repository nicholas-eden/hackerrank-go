package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the activityNotifications function below.
func activityNotifications(expenditure []int32, d int32) int32 {
	notifications := int32(0)

	counts := [200]int{}

	// Initiate counts of running window
	for _, v := range expenditure[0:d] {
		counts[int(v)]++
	}

	isEven := int(d)%2 == 0
	median := 0
	for i, v := range expenditure[d:] {

		z := 0
		for j, c := range counts {
			for ; c > 0 && z <= int(d); c-- {
				z++
				if isEven {
					median += j
				} else {
					median = j
				}
			}
		}

		if !isEven {
			median *= 2
		}

		if int(v) >= median {
			notifications++
		}

		counts[int(v)]++
		pop := expenditure[i]
		counts[pop]--
	}

	return notifications
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nd := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nd[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	dTemp, err := strconv.ParseInt(nd[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	expenditureTemp := strings.Split(readLine(reader), " ")

	var expenditure []int32

	for i := 0; i < int(n); i++ {
		expenditureItemTemp, err := strconv.ParseInt(expenditureTemp[i], 10, 64)
		checkError(err)
		expenditureItem := int32(expenditureItemTemp)
		expenditure = append(expenditure, expenditureItem)
	}

	result := activityNotifications(expenditure, d)

	fmt.Fprintf(writer, "%d\n", result)

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
