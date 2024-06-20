package leetcodego
import(
	"sort"
)
func maxDistance(position []int, m int) int {
    //credit to alanlzl, O(NlogM), where M = max(position) - min(position), Space complexity: O(1)
    sort.Ints(position)
    left, right := 0, position[len(position) - 1] - position[0]

    countNode := func(position []int, curDist int) int {
        ans, cur := 1, position[0]
        for i:= 1; i < len(position); i++{
            if position[i] - cur >= curDist{
                ans++
                cur = position[i]
            }
        }
        return ans
    }

    for left <= right{
        mid := (right - left)/2 + left
        if countNode(position, mid) >= m {
            left = mid + 1
        }else{
            right = mid - 1
        }
    }
    return right
}