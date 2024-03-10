package leetcodego

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}

 func isEqual(L *TreeNode, R *TreeNode) bool{
    if L == nil && R == nil{
        return true
    }

    if L == nil || R == nil{
        return false
    }

    if L.Val != R.Val{
        return false
    }

    return isEqual(L.Right, R.Left) && isEqual(L.Left, R.Right)
}
func isSymmetric(root *TreeNode) bool {
    return isEqual(root.Left, root.Right)
}