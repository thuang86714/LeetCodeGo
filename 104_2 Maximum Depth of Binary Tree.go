package leetcodego

func maxDepth2(root *TreeNode) int {
	//TC O(n), SC O(height of tree)
    if root == nil{
        return 0
    }
    return 1 + max(maxDepth(root.Left), maxDepth(root.Right))
}