package leetcodego

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func sortedArrayToBST(nums []int) *TreeNode {
    //credit to klakovskiy
    if len(nums) == 0{
        return nil
    }

    mid := len(nums)/2
    root := &TreeNode{Val:nums[mid]}
    root.Left = sortedArrayToBST(nums[:mid])
    root.Right = sortedArrayToBST(nums[mid + 1:])
    return root
}