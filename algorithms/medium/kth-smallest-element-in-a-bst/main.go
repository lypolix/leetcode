package main

import "fmt"


type TreeNode struct {
 	Val int
    Left *TreeNode
	Right *TreeNode
 }
func kthSmallest(root *TreeNode, k int) int {
	list := []int{}
	binsearch(root, k, &list)

	return list[k - 1]
 }
   

func binsearch(root *TreeNode, k int, list *[]int){
	if root == nil {
		return
	}
	binsearch(root.Left, k, list)

	if len(*list) == k {
		return
	}
	*list = append(*list, root.Val)

	binsearch(root.Right, k, list)
}


func main() {
	root := &TreeNode{
        Val: 3,
        Left: &TreeNode{
            Val: 1,
            Right: &TreeNode{
                Val: 2,
            },
        },
        Right: &TreeNode{
            Val: 4,
        },
    }
    k := 1
    fmt.Println(kthSmallest(root, k))
}


