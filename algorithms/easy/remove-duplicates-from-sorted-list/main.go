package main

import (
	"fmt"
)
func main() {
	node3 := &ListNode {Val : 2, Next: nil}
	node2 := &ListNode {Val : 2, Next: node3}
	node1 := &ListNode {Val : 2, Next: node2}
	

	fmt.Println(deleteDuplicates(node1))
}

type ListNode struct {
    Val int
    Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	current := head

	for current != nil && current.Next != nil {
		if current.Next.Val == current.Val {
			current.Next = current.Next.Next
		}else {
			current = current.Next
		}
	}
	
	return head 
}
