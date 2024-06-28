package leetcodego

func deleteNode(root *TreeNode, key int) *TreeNode {
    if root == nil {
        return nil
    }
    if key < root.Val {
        root.Left = deleteNode(root.Left, key)
    } else if key > root.Val {
        root.Right = deleteNode(root.Right, key)
    } else {
        if root.Left == nil && root.Right == nil {
            return nil
        }
        if root.Left == nil {
            return root.Right
        } else if root.Right == nil {
            return root.Left
        }

        temp := root.Left
        for temp.Right != nil {
            temp = temp.Right
        }
        root.Val = temp.Val
        root.Left = deleteNode(root.Left, temp.Val)
    }
    return root
}
