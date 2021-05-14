package main

import (
	"fmt"
)

func numIslands(grid [][]byte) int {
	cnt := 0
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] == 49 {
				slove(grid, i, j)
				cnt++
			}
		}
	}
	return cnt
}

func slove(grid [][]byte, i, j int) {
	if !(i >= 0 && j >= 0 && i < len(grid) && j < len(grid[0])) {
		return
	}
	if grid[i][j] == 48 {
		return
	}
	grid[i][j] = 48
	slove(grid, i+1, j)
	slove(grid, i, j+1)
	slove(grid, i-1, j)
	slove(grid, i, j-1)
}

func main() {
	// grid := [][]byte{
	// 	{49, 49, 49, 49, 48},
	// 	{49, 49, 48, 49, 48},
	// 	{49, 49, 48, 48, 48},
	// 	{48, 48, 48, 48, 48},
	// }
	grid := [][]byte{
		{49, 49, 48, 48, 48},
		{49, 49, 48, 48, 48},
		{48, 48, 49, 48, 48},
		{48, 48, 48, 49, 49},
	}
	fmt.Println(numIslands(grid))
}
