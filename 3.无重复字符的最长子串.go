package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=3 lang=golang
 *
 * [3] 无重复字符的最长子串
 */

// @lc code=start
func lengthOfLongestSubstring(s string) int {

	if len(s) <= 1 {
		return len(s)
	}
	left := 0
	maxCount := 0

	count := 0
	wordList := strings.Split(s, "")

	hasWord := map[string]bool{}

	for _, v := range wordList {
		if hasWord[v] == true {
			for hasWord[v] == true {
				delete(hasWord, wordList[left])
				left++
				count--
			}
		}

		if hasWord[v] == false {
			hasWord[v] = true
			count++

			if count > maxCount {
				maxCount = count
			}
		}
	}
	return maxCount
}

// @lc code=end

func main() {
	res := lengthOfLongestSubstring("aab")
	fmt.Println(res, res == 2)
}
