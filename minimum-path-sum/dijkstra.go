package main

import (
	"fmt"
	"math"
)

type Node struct {
	x    int
	y    int
	cost int
}

func minPathSum(intGrid [][]int) int {
	var grid [][]*Node
	q := make(map[*Node]bool)
	distQ := make(map[*Node]int)
	dist := make(map[*Node]int)
	prev := make(map[*Node]*Node)


	for y, row := range intGrid {
		var nodeRow []*Node
		for x, cost := range row {
			node := Node{x, y, cost}
			nodeRow = append(nodeRow, &node)
			q[&node] = true
		}
		grid = append(grid, nodeRow)
	}

	start := grid[0][0]
	goal := grid[len(grid)-1][ len(grid[0])-1 ]
	dist[start] = start.cost
	distQ[start] = start.cost

	var u *Node
	for ; len(q) > 0; {
		minDist := math.MaxInt64
		for node, d := range distQ {
			_, ok := q[node]
			if ok && d < minDist {
				u = node
				minDist = d
			}
		}

		delete(q, u)
		delete(distQ, u)

		if u == goal {
			return dist[goal]
		}

		ys := []int{u.y, u.y + 1}
		for _, y := range ys {
			if y > len(grid)-1 {
				continue
			}

			x := u.x
			if y == u.y {
				x = u.x + 1
			}

			row := grid[y]
			if x > len(row)-1 {
				continue
			}
			neighbor := row[x]

			if _, ok := q[neighbor]; ok == false {
				continue
			}


			d := dist[u]
			alt := d + neighbor.cost

			dn, ok := dist[neighbor]
			if !ok {
				dn = math.MaxInt64
			}
			if alt < dn {
				dist[neighbor] = alt
				distQ[neighbor] = alt
				prev[neighbor] = u
			}

		}

	}

	return dist[goal]

}

func main() {
	input := [][]int{
		{1, 3, 1},
		{1, 5, 1},
		{4, 2, 1},
	}

 	result := minPathSum(input)
	fmt.Println(result)
}
