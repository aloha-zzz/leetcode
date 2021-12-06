package main

/*
 * @lc app=leetcode.cn id=733 lang=golang
 *
 * [733] 图像渲染
 */

// @lc code=start
func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {

	prevValue := image[sr][sc]

	updateColor(image, sr, sc, prevValue, newColor)

	return image
}

func updateColor(image [][]int, row, col int, prevValue int, new int) {
	if image[row][col] == new {
		return
	}
	if image[row][col] != prevValue {
		return
	}

	image[row][col] = new
	// 上
	if row > 0 {
		updateColor(image, row-1, col, prevValue, new)
	}

	// 下
	if row < len(image)-1 {
		updateColor(image, row+1, col, prevValue, new)
	}

	// 左
	if col > 0 {
		updateColor(image, row, col-1, prevValue, new)
	}

	// 右
	if col < len(image[0])-1 {
		updateColor(image, row, col+1, prevValue, new)
	}

}

// @lc code=end

func main() {
	image := [][]int{{0, 0, 0}, {0, 1, 1}}
	floodFill(image, 1, 1, 1)

}
