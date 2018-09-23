package clb

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func binarySearch(target []int32, value int32) int32 {
	//fmt.Println("Searching for ", value)
	startIndex := 0
	endIndex := len(target) - 1

	for startIndex <= endIndex {

		median := (startIndex + endIndex) / 2

		//fmt.Println("Start ", startIndex, "End", endIndex, "Median Value", target[median])
		if target[median] > value {
			startIndex = median + 1
		} else {
			endIndex = median - 1
		}

	}

	if startIndex == len(target) {
		return int32(startIndex)
	} else if target[startIndex] != value {
		if target[startIndex] < value {
			return int32(startIndex)
		} else {
			return int32(startIndex + 1)
		}
	} else {
		return int32(startIndex)
	}

}

// Complete the climbingLeaderboard function below.
func climbingLeaderboard(scores []int32, alice []int32) []int32 {
	var uniqueScores []int32
	encountered := map[int32]bool{}

	for _, score := range scores {
		if encountered[score] == true {
			// Already in slice
		} else {
			encountered[score] = true
			uniqueScores = append(uniqueScores, score)
		}
	}

	//fmt.Println(uniqueScores)
	var positions []int32
	for _, num := range alice {
		position := binarySearch(uniqueScores, num)
		positions = append(positions, position + int32(1))
	}

	return positions
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	fmt.Println("Getting Scores Count")
	scoresCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	fmt.Println("Getting Scores")
	scoresTemp := strings.Split(readLine(reader), " ")

	var scores []int32

	for i := 0; i < int(scoresCount); i++ {
		scoresItemTemp, err := strconv.ParseInt(scoresTemp[i], 10, 64)
		checkError(err)
		scoresItem := int32(scoresItemTemp)
		scores = append(scores, scoresItem)
	}

	fmt.Println("Getting Alice Count")
	aliceCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	fmt.Println("Getting Alice Scores")
	aliceTemp := strings.Split(readLine(reader), " ")

	var alice []int32

	for i := 0; i < int(aliceCount); i++ {
		aliceItemTemp, err := strconv.ParseInt(aliceTemp[i], 10, 64)
		checkError(err)
		aliceItem := int32(aliceItemTemp)
		alice = append(alice, aliceItem)
	}

	result := climbingLeaderboard(scores, alice)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

		if i != len(result) - 1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

	writer.Flush()
}

func readLine(reader *bufio.Reader) string {
	var line string
	for {
		fmt.Println("Reading line")
		str, isPrefix, err := reader.ReadLine()
		if err == io.EOF {
			return ""
		}

		line += string(str)
		if isPrefix == false {
			break
		}
		fmt.Println("Line was big!")
	}


	return strings.TrimRight(line, "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
