package leetcodego

func isSameTree2(p *TreeNode, q *TreeNode) bool {
	//TC O(n), SC O(height of tree)
	if p == nil && q == nil {
		return true
	}
	if p == nil || q == nil {
		return false
	}
	if p.Val != q.Val {
		return false
	}
	return isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
}
