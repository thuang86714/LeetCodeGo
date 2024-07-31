package leetcodego

func invertTree2(root *TreeNode) *TreeNode {
    if root == nil{
        return nil
    }

    newL := invertTree(root.Right)
    newR := invertTree(root.Left)

    root.Left = newL
    root.Right = newR

    return root
}