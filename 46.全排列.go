package main

import "fmt"

/*
 * @lc app=leetcode.cn id=46 lang=golang
 *
 * [46] 全排列
 */

// @lc code=start
func permute(nums []int) [][]int {

	usedMap := map[int]bool{}
	resultArr := [][]int{}
	for k, v := range nums {
		backTrace(nums, usedMap, &resultArr, []int{}, k, v)
	}
	return resultArr
}

// DFS 和回溯, 可以认为回溯实在 DFS 后面加了逻辑
func backTrace(nums []int, usedMap map[int]bool, resultArr *[][]int, tmpValue []int, k, v int) {

	usedMap[k] = true
	tmpValue = append(tmpValue, v)

	if len(tmpValue) == len(nums) {
		*resultArr = append(*resultArr, tmpValue)
		// 回溯
		usedMap[k] = false
		removeLast(tmpValue)
		return
	}

	for idx, value := range nums {
		if usedMap[idx] {
			continue
		}
		backTrace(nums, usedMap, resultArr, tmpValue, idx, value)
	}
	usedMap[k] = false
	removeLast(tmpValue)
}

func removeLast(arr []int) []int {
	length := len(arr)
	return arr[:length-1]

}

// @lc code=end

func main() {
	nums := []int{1, 2, 3}
	res := permute(nums) /**/

	fmt.Println(res)
}

func operate(a *[][]int) {
	*a = append(*a, []int{1, 2, 3})
}
