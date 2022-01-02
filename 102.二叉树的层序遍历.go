package main

import "fmt"

/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a binary tree node.

 */

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := [][]int{}

	tmpList := []*TreeNode{}

	tmpList = append(tmpList, root)

	for len(tmpList) != 0 {
		tmpValue := []int{}
		newList := []*TreeNode{}
		for _, v := range tmpList {
			tmpValue = append(tmpValue, v.Val)
			if v.Left != nil {
				newList = append(newList, v.Left)
			}
			if v.Right != nil {
				newList = append(newList, v.Right)
			}
		}
		tmpList = newList
		res = append(res, tmpValue)
	}

	return res
}

// @lc code=end

func main() {
	a := TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
		},
	}

	res := levelOrder(&a)
	fmt.Println(res)
}
