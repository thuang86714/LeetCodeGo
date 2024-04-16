package leetcodego

func pathSum(root *TreeNode, targetSum int) int {
    if root == nil {
        return 0
    }
    
    var findAllValidPath func(node *TreeNode, curSum int64) int
    findAllValidPath = func(node *TreeNode, curSum int64) int {
        if node == nil {
            return 0
        }
        // Current path count is 0
        currentPathCount := 0
        
        // Check if current node completes a valid path
        if int64(node.Val) == curSum {
            currentPathCount = 1
        }
        
        // Add counts from left and right children
        currentPathCount += findAllValidPath(node.Left, curSum - int64(node.Val))
        currentPathCount += findAllValidPath(node.Right, curSum - int64(node.Val))
        
        return currentPathCount
    }

    // Check paths starting from the current root
    pathCount := findAllValidPath(root, int64(targetSum))
    
    // Recursively count paths from left and right children
    pathCount += pathSum(root.Left, targetSum)
    pathCount += pathSum(root.Right, targetSum)
    
    return pathCount
}
