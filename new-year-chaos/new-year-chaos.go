package nyc

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the minimumBribes function below.
func minimumBribes(q []int32) {

	var bribeCount int
	var currentBribeCount int
	for i, num := range q {

		currentBribeCount = 0
		for j := i + 1; j < i + 4 && j < len(q); j++ {
			if q[j] < num {
				bribeCount += 1
				currentBribeCount += 1
			}

			fmt.Println(i, num, j, q[j], currentBribeCount, bribeCount)

			if currentBribeCount > 2 {
				fmt.Println("Too chaotic")
				return
			}

		}

	}

	fmt.Println(bribeCount)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		qTemp := strings.Split(readLine(reader), " ")

		q := make([]int32, int(n))

		for i := 0; i < int(n); i++ {
			qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
			checkError(err)
			q[i] = int32(qItemTemp)
		}

		minimumBribes(q)
	}
}

func readLine(reader *bufio.Reader) string {
	var line string
	for {
		str, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			return ""
		}

		line += string(str)
		if isPrefix == false {
			break
		}
	}


	return strings.TrimRight(line, "\r\n")
}


func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
