package main

import "fmt"

func numIslands(grid [][]byte) int {
	CNT := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			//
			if grid[i][j] == byte(1) {
				CNT++
			}
		}
	}
	return CNT
}

func main() {
	grid := [][]byte{
		{1, 1, 1, 1, 0},
		{1, 1, 0, 1, 0},
		{1, 1, 0, 0, 0},
		{0, 0, 0, 0, 0}}

	res := numIslands(grid)
	fmt.Println(res)
}
