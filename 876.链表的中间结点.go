package leetcode

import "fmt"

/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	count := 0

	left := head
	for left != nil {
		left = left.Next
		count++
	}

	if count == 1 {
		return head
	}
	fmt.Println(count)
	count = count / 2
	fmt.Println(count)
	right := head

	for count > 0 {
		right = right.Next
		fmt.Println(right.Val)
		count--
	}
	return right
}

// @lc code=end

func main() {

}
