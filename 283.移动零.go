package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=283 lang=golang
 *
 * [283] 移动零
 */

// @lc code=start
func moveZeroes(nums []int) {
	length := len(nums)
	if length <= 1 {
		return
	}
	for i := length - 1; i >= 0; i-- {
		if nums[i] == 0 {
			swapRight(nums, i)
		}
	}

}

func findFirstNonZeroIdx(arr []int) int {
	for k, v := range arr {
		if v != 0 {
			return k
		}
	}
	return -1
}

func swapRight(arr []int, idx int) {
	if idx == len(arr)-1 {
		return
	}

	tmp := arr[idx]

	arr = append(arr[:idx], arr[idx+1:]...)
	arr = append(arr, tmp)

}

// @lc code=end

func main() {
	a := []int{0, 1, 0, 3, 12}

	moveZeroes(a)

	fmt.Println(a)

}
