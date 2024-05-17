package leetcodego

func evaluateTree(root *TreeNode) bool {
    //TC O(n), SC O(height of tree)
    if root == nil{
        return false
    }

    if root.Val > 1{
        if root.Val == 2{
            return evaluateTree(root.Left) || evaluateTree(root.Right)
        }
        return evaluateTree(root.Left) && evaluateTree(root.Right)
    }
    return root.Val > 0
}