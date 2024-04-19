package leetcodego

func smallestFromLeaf(root *TreeNode) string {
    //credit to eduard1998n, TC O(n), SC O()
    ans := ""
    var findSmallestString func(node *TreeNode, curString string)
    findSmallestString = func(node *TreeNode, curString string){
        if node == nil{
            return
        }

        curString = string(rune(node.Val + 97)) + curString//append to the front, no need reverse()
        if node.Right == nil && node.Left == nil{
            if ans == "" || ans > curString{
                ans = curString
            }
            return
        }
        findSmallestString(node.Left, curString)
        findSmallestString(node.Right, curString)
    }
    findSmallestString(root, "")
    return ans
}