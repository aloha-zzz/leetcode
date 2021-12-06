package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TODO:fail
func getDirections(root *TreeNode, startValue int, destValue int) string {
	//leftPath := ""
	//rightPath := ""
	return ""
}

type level struct {
	prev string
	node *TreeNode
}

func getNode(root *TreeNode, target int, prev *string) {
	if root.Val == target {
		return
	}

	l := []*level{}
	if root.Left != nil {
		*prev += "L"
		l = append(l, &level{node: root.Left, prev: *prev})
	}

	if root.Right != nil {
		*prev += "R"
		l = append(l, &level{node: root.Left, prev: *prev})
	}

	res := BFS(l, target)
	fmt.Println(res)
}

func BFS(l []*level, target int) string {
	newList := []*level{}
	for _, v := range l {
		if v.node.Val == target {
			return v.prev
		}

		if v.node.Left != nil {
			v.prev += "L"
			newList = append(newList, &level{node: v.node.Left, prev: v.prev})
		}

		if v.node.Right != nil {
			v.prev += "R"
			newList = append(newList, &level{node: v.node.Right, prev: v.prev})
		}
	}
	return BFS(newList, target)
}

func main() {
	tree := &TreeNode{Val: 5, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 4}}, Left: &TreeNode{Val: 1, Left: &TreeNode{Val: 3}}}
	str := ""
	getNode(tree, 3, &str)
	getNode(tree, 6, &str)

}
