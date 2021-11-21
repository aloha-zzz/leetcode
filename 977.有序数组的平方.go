package main

import "fmt"

/*
 * @lc app=leetcode.cn id=977 lang=golang
 *
 * [977] 有序数组的平方
 */
// https://code-thinking.cdn.bcebos.com/gifs/977.%E6%9C%89%E5%BA%8F%E6%95%B0%E7%BB%84%E7%9A%84%E5%B9%B3%E6%96%B9.gif
/**
其实考虑复杂了， 只要在两端进行双指针即可。
*/
// 1 找 最小值
// @lc code=start
func sortedSquares(nums []int) []int {
	length := len(nums)
	res := []int{}
	if length <= 1 {
		for _, v := range nums {
			res = append(res, v*v)
		}
		return res
	}

	idx := 0
	for k, v := range nums {
		if v >= 0 {
			idx = k
			break
		} else {
			idx++
		}
	}

	if idx > 0 {
		idx--
	}
	fmt.Println(idx)
	// 第一个负数是 idx

	left := idx
	right := idx + 1

	for left >= 0 && right <= length-1 {
		a := nums[left] * nums[left]
		b := nums[right] * nums[right]

		if a > b {
			res = append(res, b)
			right++
		} else {
			res = append(res, a)
			left--
		}
	}

	if left >= 0 {
		for i := left; i >= 0; i-- {
			res = append(res, nums[i]*nums[i])
		}
	}

	if right <= length-1 {
		for i := right; i < length; i++ {
			res = append(res, nums[i]*nums[i])
		}
	}

	return res
}

// @lc code=end
func main() {
	input := []int{-5, -3, -2, -1}

	res := sortedSquares(input)
	fmt.Printf("res %+v", res)
}
