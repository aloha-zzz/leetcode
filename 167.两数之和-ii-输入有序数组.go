package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=167 lang=golang
 *
 * [167] 两数之和 II - 输入有序数组
 */

// 双指针搜索的原理
//  * https://leetcode-cn.com/problems/two-sum-ii-input-array-is-sorted/solution/yi-zhang-tu-gao-su-ni-on-de-shuang-zhi-zhen-jie-fa/

// @lc code=start
func twoSum(numbers []int, target int) []int {
	if len(numbers) <= 1 {
		return nil
	}
	length := len(numbers)
	left := 0

	right := length - 1

	for numbers[left]+numbers[right] != target {
		if numbers[left]+numbers[right] > target {
			right--
		} else {
			left++
		}
	}
	return []int{left + 1, right + 1}
}

// @lc code=end

func main() {
	a := []int{2, 7, 11, 15}
	target := 9
	res := twoSum(a, target)
	fmt.Println(res)
}
