package main

import "fmt"

/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.

 */

// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }

// 1 -> 2 - >3
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var prev, tmp *ListNode
	cur := head

	for cur != nil {
		tmp = cur

		// 进行下一步
		cur = cur.Next
		tmp.Next = prev
		prev = tmp
	}

	return prev
}

// @lc code=end

func (l *ListNode) Append(num int) {
	if l == nil {
		l = &ListNode{Val: num}
		return
	}

	h := l
	for h.Next != nil {
		h = h.Next
	}

	h.Next = &ListNode{Val: num}
}
func (l *ListNode) Output() {
	h := l
	if h == nil {
		return
	}

	for h != nil {
		fmt.Println(h.Val)
		h = h.Next
	}
}

func main() {
	a := ListNode{Val: 1}

	a.Append(2)

	res := reverseList(&a)
	res.Output()
}

/**
// 通过栈
func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	stack := []int{}

	for head.Next != nil {
		stack = append(stack, head.Val)
		head = head.Next
	}
	last := head

	for i := len(stack) - 1; i >= 0; i-- {
		last.Next = &ListNode{Val: stack[i]}
		last = last.Next
	}

	return head
}

*/
