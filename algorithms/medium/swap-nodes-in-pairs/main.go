package main

import "fmt"

func main() {
    head := createList([]int{1, 2, 3, 4})
    swapped := swapPairs(head)
    printList(swapped)
}

func createList(arr []int) *ListNode {
    dummy := &ListNode{}
    cur := dummy
    for _, v := range arr {
        cur.Next = &ListNode{Val: v}
        cur = cur.Next
    }
    return dummy.Next
}

func printList(head *ListNode) {
    cur := head
    for cur != nil {
        fmt.Print(cur.Val)
        cur = cur.Next
    }
}


type ListNode struct {
    Val int
    Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {

	dummy := &ListNode{Next: head}
	prev := dummy

	for prev.Next != nil && prev.Next.Next != nil {
		first := prev.Next
		second := prev.Next.Next

		prev.Next = second
		first.Next = second.Next
		second.Next = first
		
		prev = first
	}
	return dummy.Next
  	
}
