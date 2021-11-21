package main

import "fmt"

/*
 * @lc app=leetcode.cn id=35 lang=golang
 *
 * [35] 搜索插入位置
 */

// @lc code=start
func searchInsert(nums []int, target int) int {
	left := 0
	right := len(nums) - 1

	for {
		mid := (left + right) / 2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] > target {
			right = mid - 1
		}

		if nums[mid] < target {
			left = mid + 1
		}
		fmt.Println(left, right, mid)
		if right <= left {
			if right < 0 {
				right = 0
			}

			if nums[right] >= target {
				return right
			}
			return right + 1
		}
	}

}

// @lc code=end

func main() {
	num := []int{1, 3}
	target := 3

	res := searchInsert(num, target)
	fmt.Println(res)
}
