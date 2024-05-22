package leetcodego

var maxStep int;

func findMaxLength(root *TreeNode, isLeft bool, step int){
    if root == nil{
        return
    }

    maxStep = max(maxStep, step)

    if isLeft{
        findMaxLength(root.Left, false, step + 1)// keep going from root to left
        findMaxLength(root.Right, true, 1)// restart going from root to right
    }else{
        findMaxLength(root.Right, true, step + 1)// keep going from root to left
        findMaxLength(root.Left, false, 1)// restart going from root to right
    }
}

func longestZigZag(root *TreeNode) int {
    maxStep = 0
    findMaxLength(root, true, 0)
    findMaxLength(root, false, 0)
    return maxStep
}