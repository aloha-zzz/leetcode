/*
 * @lc app=leetcode.cn id=344 lang=golang
 *
 * [344] 反转字符串
 */

// @lc code=start
func reverseString(s []byte) {
	left := 0
	right := len(s) - 1

	for right >= left {
		swap(s, left, right)
		left++
		right--
	}
}

func swap(arr []byte, a, b int) {
	arr[a], arr[b] = arr[b], arr[a]
}

// @lc code=end

