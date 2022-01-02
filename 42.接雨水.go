package main

import "fmt"

/*
 * @lc app=leetcode.cn id=42 lang=golang
 *
 * [42] 接雨水
 */

// @lc code=start
func trap(height []int) int {

	midMax := 0
	midIdx := -1

	for k, v := range height {
		if v > midMax {
			midIdx = k
			midMax = v
		}
	}

	sum := 0

	leftMax := 0

	for i := 0; i < midIdx; i++ {
		if leftMax < height[i] {
			leftMax = height[i]
		} else {
			sum += leftMax - height[i]
		}
	}

	rightMax := 0
	for i := len(height) - 1; i > midIdx; i-- {
		if rightMax < height[i] {
			rightMax = height[i]
		} else {
			sum += rightMax - height[i]
		}
	}
	return sum
}

// @lc code=end

func main() {
	tmp := []int{4, 2, 0, 3, 2, 5}
	sum := trap(tmp)
	fmt.Println(sum)
}
