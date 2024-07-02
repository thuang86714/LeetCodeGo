package leetcodego
func containsNearbyDuplicate(nums []int, k int) bool {
    //TC O(n), SC O(n)
    indexMap := make(map[int]int)//using map to keep track of last visited index number
    for idx, num := range nums{
        if val, ok := indexMap[num]; ok{
            if idx - val <= k{
                return true
            }
        }
        indexMap[num] = idx
    }
    return false
}