package leetcodego

import(
	"sort"
)

func maximumHappinessSum(happiness []int, k int) int64 {
    //TC O(nlogn), SC O(1)
    var ans int64
    childCount, idx := 0, len(happiness) - 1
    sort.Ints(happiness)
    for k > 0 && idx >= 0{
        if happiness[idx] - childCount >= 0{
            ans += int64(happiness[idx] - childCount)
        }else{//meaning that we've reached the element that woudl give us max(given k)
            break
        }
        idx--
        k--
        childCount++
    }
    return ans
}