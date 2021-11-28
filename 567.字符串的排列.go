package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=567 lang=golang
 *
 * [567] 字符串的排列
 */

// @lc code=start
func checkInclusion(s1 string, s2 string) bool {
	length := len(s1)

	w1 := strings.Split(s1, "")
	w2 := strings.Split(s2, "")

	wordMapTarget := map[string]int{}

	for _, v := range w1 {
		wordMapTarget[v] += 1
	}

	left := 0
	right := length - 1

	if len(s1) > len(s2) {
		return false
	}

	wordMap := map[string]int{}
	for k, v := range w2 {
		if k < right {
			wordMap[v] += 1
			continue
		}
		wordMap[v] += 1
		fmt.Println(wordMap)
		if isEqualMap(wordMapTarget, wordMap) {
			return true
		}
		wordMap[w2[left]] -= 1 // 删除前一个

		left++

	}
	return false
}

func isEqualMap(s1, s2 map[string]int) bool {
	for k, v := range s1 {
		if v != s2[k] {
			return false
		}
	}
	return true
}

// @lc code=end
func main() {
	s1 := "ab"
	s2 := "eidbaooo"

	res := checkInclusion(s1, s2)

	fmt.Println(res)

}
