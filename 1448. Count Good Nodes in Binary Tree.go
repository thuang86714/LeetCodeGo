package leetcodego

func goodNodes(root *TreeNode) int {
    //TC O(n), SC O(logn)
    goodNodeCount := 0
    var findGoodNode func(*TreeNode, int)

    findGoodNode = func(node *TreeNode, curMax int){
        if node == nil{
            return
        }

        if node.Val >= curMax{
            goodNodeCount++
            curMax = node.Val
        }

        findGoodNode(node.Left, curMax)
        findGoodNode(node.Right, curMax)
    }
    findGoodNode(root, -1e5)
    return goodNodeCount 
}