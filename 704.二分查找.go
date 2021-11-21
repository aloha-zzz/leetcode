package main

import "fmt"

/*
 * @lc app=leetcode.cn id=704 lang=golang
 *
 * [704] 二分查找
 */

// @lc code=start
func search(nums []int, target int) int {
	length := len(nums)
	if len(nums) == 0 {
		return -1
	}

	left := 0
	right := length - 1

	for left <= right {
		mid := (right + left) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
	}
	return -1
}

// @lc code=end
func main() {
	nums := []int{-1, 0, 3, 5, 9, 12}
	target := 9

	res := search(nums, target)

	fmt.Printf("res %+v", res)
}
