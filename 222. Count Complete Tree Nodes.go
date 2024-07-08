package leetcodego

 func countNodes(root *TreeNode) int {
    leftHeight, rightHeight:= 0, 0
    curR, curL := root, root
    for curL != nil{
        curL = curL.Left
        leftHeight++
    }
    for curR != nil{
        curR = curR.Right
        rightHeight++
    }

    if leftHeight == rightHeight{
        return (1 << leftHeight) - 1
    }else{
        return 1 + countNodes(root.Left) + countNodes(root.Right)
    }
}