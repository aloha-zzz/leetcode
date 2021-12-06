package main

import "fmt"

/*
 * @lc app=leetcode.cn id=695 lang=golang
 *
 * [695] 岛屿的最大面积
 */

// @lc code=start
func maxAreaOfIsland(grid [][]int) int {
	area := 0
	row := 0

	for ; row < len(grid); row++ {

		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 1 {

				i := calcArea(grid, row, col)
				if i > area {
					area = i
				}
			}
		}
	}
	return area
}

func calcArea(grid [][]int, row, col int) int {
	val := 0

	DFS(grid, row, col, &val)

	return val
}

func DFS(grid [][]int, row, col int, prev *int) {

	if grid[row][col] != 1 {
		return
	}
	grid[row][col] = 0
	add(prev)
	// 上
	if row > 0 {
		DFS(grid, row-1, col, prev)
	}
	// 下
	if row < len(grid)-1 {
		DFS(grid, row+1, col, prev)
	}
	//左
	if col > 0 {
		DFS(grid, row, col-1, prev)
	}

	// 右

	if col < len(grid[0])-1 {
		DFS(grid, row, col+1, prev)
	}
}

func add(a *int) {
	tmp := *a
	tmp++
	*a = tmp
}

// @lc code=end

func main() {

	tmp := [][]int{{0, 0, 1, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 1, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 0, 1, 0, 0}, {0, 1, 0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 0, 0, 0}, {0, 0, 0, 0, 0, 0, 0, 1, 1, 0, 0, 0, 0}}
	res := maxAreaOfIsland(tmp)

	fmt.Println(res)
}
