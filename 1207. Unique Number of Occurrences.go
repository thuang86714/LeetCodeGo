package leetcodego
func uniqueOccurrences(arr []int) bool {
    //credit to miryang, two unordered_map, 
    countTracker := make(map[int]int)
    for _, num:=range(arr){
        countTracker[num]++
    }
    result := make(map[int]int)
    for _,val := range countTracker{
        result[val]++
        if result[val] != 1{
            return false
        }
    }
    return true
}