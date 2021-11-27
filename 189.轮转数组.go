package main

import "fmt"

/*
 * @lc app=leetcode.cn id=189 lang=golang
 *
 * [189] 轮转数组
 */
 // 可以通过 pop  unshift 的 方法 操作数据

// @lc code=start
func rotate(nums []int, k int) {

	length := len(nums)

	k = k % length
	if k == 0 {
		return
	}

	res := []int{}

	res = append(res, nums[length-k:]...)

	res = append(res, nums[:length-k]...)

	for k := range nums {
		nums[k] = res[k]
	}
}

// @lc code=end
func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7}

	k := 3
	rotate(nums, k)

	fmt.Printf("Res %+v", nums)
}
