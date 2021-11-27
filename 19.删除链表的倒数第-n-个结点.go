package main

import (
	"fmt"
)

/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	tmp := n

	slow := head
	fast := head

	prev := slow

	for fast.Next != nil {
		fast = fast.Next
		fmt.Println("fast", fast.Val)
		tmp--
		if tmp <= 0 {
			prev = slow
			slow = slow.Next
			fmt.Println("slow", slow.Val)
		}
	}
	fmt.Println(tmp)
	if tmp > 0 {
		if prev.Next == nil {
			return nil
		}
		// 此时代表要删除头
		prev = prev.Next
		return prev
	}

	// 此时 f

	prev.Next = prev.Next.Next
	return head
}

// @lc code=end

func main() {
	// case 3
	z := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
			Next: &ListNode{
				Val: 3,
				Next: &ListNode{
					Val: 4,
					Next: &ListNode{
						Val: 5,
					},
				},
			},
		},
	}
	res := removeNthFromEnd(z, 2)
	// TODO: add method
	fmt.Println(res.Val, res.Next.Val, res.Next.Next.Val, res.Next.Next.Next.Val, res.Next.Next.Next.Val == 5)

	// case 1
	v := &ListNode{
		Val: 1,
	}
	res = removeNthFromEnd(v, 1)
	fmt.Println((res))

	// // case 2

	v2 := &ListNode{
		Val: 1,
		Next: &ListNode{
			Val: 2,
		},
	}
	fmt.Println("case 3")

	res = removeNthFromEnd(v2, 2)
	fmt.Println(res)

}
