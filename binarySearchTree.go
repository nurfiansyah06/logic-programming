package logicprogramming

type TreeNode struct {
	Val int
	Left *TreeNode
	Right *TreeNode
}	

func sortedArrayToBST(nums []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}

    i := len(nums)/2
    root := &TreeNode{Val: nums[i]}
    root.Left = sortedArrayToBST(nums[:i])
    root.Right = sortedArrayToBST(nums[i+1:])
    return root
}