package leetcodego

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