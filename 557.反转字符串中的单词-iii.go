package main

import (
	"fmt"
	"strings"
)

/*
 * @lc app=leetcode.cn id=557 lang=golang
 *
 * [557] 反转字符串中的单词 III
 */

// @lc code=start
func reverseWords(s string) string {
	list := strings.Split(s, " ")

	for k, v := range list {
		tmp := []byte(v)
		fmt.Println("tmp v", tmp, v)
		reverseString(tmp)
		fmt.Println("tmp 2 str ", tmp, string(tmp))
		list[k] = string(tmp)
	}
	fmt.Println(list)
	return strings.Join(list, " ")
}
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

func main() {
	a := "Let's take LeetCode contest"
	res := reverseWords(a)
	fmt.Println(res)

	l := []int{1, 2, 3}
	for _, v := range l {
		fmt.Println(v)
		v = 6
	}

	fmt.Println(l)

}
