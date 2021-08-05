package main

import "testing"

func TestAddTwoNumbers(t *testing.T)  {
	l1 := new(ListNode)
	l2 := new(ListNode)

	l1LastNode := l1
	l2LastNode := l2

	for i := 0; i < 10; i++ {
		l1LastNode.Val = i
		l1TmpNode := new(ListNode)
		l1LastNode.Next = l1TmpNode
		l1LastNode = l1TmpNode

		l2LastNode.Val = i
		l2TmpNode := new(ListNode)
		l2LastNode.Next = l2TmpNode
		l2LastNode = l2TmpNode
	}

	addTwoNumbers(l1, l2)
}

