package main

import "fmt"

func main() {
	input := [][]byte{
		{'1', '1', '1'},
		{'0', '1', '0'},
		{'1', '1', '1'},
	}

	result := numIslands(input)
	fmt.Println(result)
}

func numIslands(grid [][]byte) int {
	islands := make([][]int, len(grid))

	// Set up matrix size
	for y, row := range grid {
		for x := 0; x < len(row); x++ {
			islands[y] = append(islands[y], 0)
		}
	}

	count := 0
	for y, row := range grid {
		for x, v := range row {
			if v == byte(48) {
				continue
			}

			if islands[y][x] > 0 {
				continue
			}

			count = count + 1

			traverse(&islands, &grid, x, y, count)
		}
	}


	return count
}

func traverse(islands *[][]int, grid *[][]byte, x int, y int, id int)  {
	if (*grid)[y][x] == byte(48) {
		// not applicable
		return
	}

	if (*islands)[y][x] > 0 {
		// already visited
		return
	}

	(*islands)[y][x] = id
	width := len((*islands)[0])

	if y > 0 {
		traverse(islands, grid, x, y - 1, id)
	}

	if y < len(*islands) - 1 {
		traverse(islands, grid, x, y + 1, id)
	}

	if x > 0 {
		traverse(islands, grid, x - 1, y, id)
	}

	if x < width - 1 {
		traverse(islands, grid, x + 1, y, id)
	}

}