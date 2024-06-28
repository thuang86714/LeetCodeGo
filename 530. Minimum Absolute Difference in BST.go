package leetcodego
import(
    "math"
)
//credit to harotobira
var curMin = math.MaxInt32
func getMinimumDifference(root *TreeNode) int {
    curMin := math.MaxInt64
    preVal := math.MinInt32

    findDiff := func(n *TreeNode){}
    findDiff = func(n *TreeNode){
        if n == nil{
            return
        }
        findDiff(n.Left)
        if preVal != math.MinInt32{
            d := n.Val - preVal
            if curMin > d{
                curMin = d
            }
        }

        preVal = n.Val
        findDiff(n.Right)
    }

    findDiff(root)
    return curMin
}