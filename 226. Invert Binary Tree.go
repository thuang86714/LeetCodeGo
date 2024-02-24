package leetcodedynamicprogramming

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func invertTree(root *TreeNode) *TreeNode {
    if root == nil{
        return root
    }

    newL := invertTree(root.Right)
    newR := invertTree(root.Left)

    root.Left = newL
    root.Right = newR

    return root
}