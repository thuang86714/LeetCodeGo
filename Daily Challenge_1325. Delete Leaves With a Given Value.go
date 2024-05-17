package leetcodego

func findLeaves(node *TreeNode, target int) *TreeNode{
    if node == nil{
        return nil
    }

    node.Left = findLeaves(node.Left, target)
    node.Right = findLeaves(node.Right, target)

    //meaning this is the target leaf
    if node.Left == nil && node.Right == nil && node.Val == target{
        return nil
    }
    return node
}
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
    //TC O(n), SC O(height of tree)
    return findLeaves(root, target)
}