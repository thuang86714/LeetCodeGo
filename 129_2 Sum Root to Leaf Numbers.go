package leetcodego

func findSum2(node *TreeNode, sum int) int{
    ans := 0
    sum = sum*10 + node.Val
    if node.Left ==nil && node.Right == nil{
        return sum
    }
    
    if node.Left != nil{
        ans += findSum(node.Left, sum)
    }
    if node.Right != nil{
        ans += findSum(node.Right, sum)
    }
    return ans
}
func sumNumbers2(root *TreeNode) int {
    return findSum(root, 0)
}