package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteMiddle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	idx := 1

	tmp := head
	for tmp != nil {
		tmp = tmp.Next
		idx++
	}
	if idx == 2 {
		return nil
	}

	removeNode := 0
	if idx%2 == 0 {
		removeNode = idx / 2
	} else {
		removeNode = idx/2 + 1
	}
	fmt.Println(removeNode)
	prev := head

	a := head

	for removeNode > 1 {
		prev = a
		a = a.Next
		removeNode--
	}

	prev.Next = a.Next

	return head
}

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
		fmt.Print(h.Val)
		h = h.Next
	}
	fmt.Println("-output-")
}
func main() {
	head := ListNode{Val: 1}

	res := deleteMiddle(&head)
	res.Output()

	head2 := ListNode{Val: 1}
	head2.Append(3)
	head2.Append(4)

	head2.Append(7)

	head2.Append(1)
	head2.Append(2)
	head2.Append(6)
	res2 := deleteMiddle(&head2)
	res2.Output()
}
