package main

import "fmt"

/*
 * @lc app=leetcode.cn id=278 lang=golang
 *
 * [278] 第一个错误的版本
 */

// @lc code=start
/**
 * Forward declaration of isBadVersion API.
 * @param   version   your guess about first bad version
 * @return 	 	      true if current version is bad
 *			          false if current version is good
 * func isBadVersion(version int) bool;
 */

func firstBadVersion(n int) int {

	left := 0
	right := n

	// case1 := false
	// case2 := false

	for {
		mid := (left + right) / 2
		if isBadVersion(mid) == true {
			right = mid

		}
		if isBadVersion(mid) == false {
			left = mid

		}

		fmt.Println(left, mid, right)
		if right-left <= 1 {
			if isBadVersion(left) {
				return left
			}
			return right
		}
	}

}

// @lc code=end
func isBadVersion(version int) bool {

	if version >= 4 {
		return true
	}
	return false
}
func main() {
	res := firstBadVersion(5)

	fmt.Printf("res %+v", res)
}
