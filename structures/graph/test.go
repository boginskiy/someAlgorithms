package main

import "fmt"

func bfs3(grid [][]byte, row, col int) {
	grid[row][col] = '#'
	queue := [][]int{{row, col}}

	// Границы
	maxRow, maxCol := len(grid), len(grid[0])

	for len(queue) > 0 {
		currentP := queue[0]
		queue = queue[1:]

		for _, shift := range [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}} {
			newRow := currentP[0] + shift[0]
			newCol := currentP[1] + shift[1]

			if newRow >= 0 && newRow < maxRow &&
				newCol >= 0 && newCol < maxCol &&
				grid[newRow][newCol] == '1' {

				grid[newRow][newCol] = '#'
				queue = append(queue, []int{newRow, newCol})
			}
		}
	}
}

func numIslands(grid [][]byte) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	rows, cols := len(grid), len(grid[0])
	cnt := 0

	for row := 0; row < rows; row++ {
		for col := 0; col < cols; col++ {
			if grid[row][col] == '1' {
				bfs3(grid, row, col)
				cnt++
			}
		}
	}
	return cnt
}

func main() {
	grid := [][]byte{
		{'1', '1', '1', '1', '0'},
		{'1', '1', '0', '1', '0'},
		{'1', '1', '0', '0', '0'},
		{'0', '0', '0', '0', '0'}}

	res := numIslands(grid)
	fmt.Println(res)
}
