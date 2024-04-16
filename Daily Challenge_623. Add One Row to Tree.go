package leetcodego

func addNewNode(node *TreeNode, val, depth, curDepth int){
    if node == nil{
        return
    }
    if curDepth < depth - 1{
        addNewNode(node.Left, val, depth, curDepth + 1)
        addNewNode(node.Right, val, depth, curDepth + 1)
        return
    }

    oldLeftPtr, oldRightPtr := node.Left, node.Right
    newLeftPtr, newRightPtr := &TreeNode{Val:val}, &TreeNode{Val:val}
    newLeftPtr.Left = oldLeftPtr
    newRightPtr.Right = oldRightPtr
    node.Left = newLeftPtr
    node.Right = newRightPtr
    return
}
func addOneRow(root *TreeNode, val int, depth int) *TreeNode {
	/*
	The function would be called for every node in the binary tree up to depth - 1. If the tree is a complete binary tree, each level doubles the number of nodes from the previous level. 
	The number of calls, therefore, in the worst case, grows exponentially with the height of the tree.
	TC O(2^(d - 1)), SC O(d)
	*/
	
    if depth == 1{
        newRoot := TreeNode{Val: val}
        newRoot.Left = root
        return &newRoot
    }
    addNewNode(root, val, depth, 1)
    return root
}