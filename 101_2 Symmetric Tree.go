package leetcodego

//TC O(n), SC O(height of tree)
func isEqual2(right, left *TreeNode) bool{
    if right == nil && left == nil{
        return true
    }
    if right == nil || left == nil{
        return false
    }
    if right.Val != left.Val{
        return false
    }
    return isEqual(right.Right, left.Left) && isEqual(right.Left, left.Right)
}
func isSymmetric2(root *TreeNode) bool {
    return isEqual(root.Right, root.Left)
}