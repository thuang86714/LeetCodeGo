package leetcodego

 //TC O(n), SC O(height of tree)
 func collectSum(node *TreeNode, cur int) int {
    if node == nil {
        return cur
    }
    node.Val += collectSum(node.Right, cur)
    return collectSum(node.Left, node.Val)
}
func bstToGst(root *TreeNode) *TreeNode {
    collectSum(root, 0)
    return root
}