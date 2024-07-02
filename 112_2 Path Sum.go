package leetcodego


 func hasPathSum2(root *TreeNode, targetSum int) bool {
    if root == nil{
        return false
    }
    targetSum -= root.Val
    if targetSum == 0 && root.Left == nil && root.Right == nil{
        return true
    }

    return hasPathSum(root.Left, targetSum) || hasPathSum(root.Right, targetSum)
}