package leetcodego

//  func findBottomLeftValue(root *TreeNode) int {
//     //BFS, TC O(n), SC O(n)
// 	if root == nil{
// 		return -1
// 	}
//     treeNodePtrSlice := []*TreeNode{root}//there's no queue in Go, we use slice instead
//     var ans int
//     for len(treeNodePtrSlice) != 0{
//         curSize := len(treeNodePtrSlice)
//         for i:= 0; i < curSize; i++{
//             cur := treeNodePtrSlice[0]
//             treeNodePtrSlice = treeNodePtrSlice[1:]
//             ans = cur.Val
//             if(cur.Right != nil){
//                 treeNodePtrSlice = append(treeNodePtrSlice, cur.Right)
//             }
//             if(cur.Left != nil){
//                 treeNodePtrSlice = append(treeNodePtrSlice, cur.Left)
//             }
//         }
//     }
//     return ans
// }

// func findBottomLeftValue(root *TreeNode) int {
//     //DFS, TC O(n), SC O(logn)
// 	curMaxDepth := 0
//     ans := 0
//     var findBottomLeftDfs func(*TreeNode, int)
//     findBottomLeftDfs = func(node *TreeNode, curDepth int){
//         if node == nil{
//             return
//         }
//         if ans[0] <= curDepth{
//             ans[0] = curDepth
//             ans[1] = node.Val
//         }
//         findBottomLeftDfs(node.Right, curDepth + 1)
//         findBottomLeftDfs(node.Left, curDepth + 1)
//     }
//     findBottomLeftDfs(root, 0)
//     return ans[1]
// }

func findBottomLeftValue(root *TreeNode) int {
    //DFS, TC O(n), SC O(logn)
    ans := 0
    curMaxDepth := 0
    var findBottomLeftDfs func(*TreeNode, int)
    findBottomLeftDfs = func(node *TreeNode, curDepth int){
        if node == nil{
            return
        }
        if curMaxDepth <= curDepth{
            curMaxDepth = curDepth
            ans = node.Val
        }
        findBottomLeftDfs(node.Right, curDepth + 1)
        findBottomLeftDfs(node.Left, curDepth + 1)
    }
    findBottomLeftDfs(root, 0)
    return ans
}