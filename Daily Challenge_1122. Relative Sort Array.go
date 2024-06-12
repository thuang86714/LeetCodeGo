package leetcodego
import(
	"sort"
)
func relativeSortArray(arr1 []int, arr2 []int) []int {
    //TC O(nlogn), SC O(n)
    numCountMap := make(map[int]int)
    for _, num := range arr1{
        numCountMap[num]++
    }

    idx := 0
    for _, num := range arr2{
        for numCountMap[num] > 0{
            arr1[idx] = num
            idx++
            numCountMap[num]--
        }
    }

    restSlice := []int{}
    for key, value := range numCountMap{
        for value > 0{
            restSlice = append(restSlice, key)
            value--
        }
    }
    sort.Ints(restSlice)
    arr1 = append(arr1[:idx], restSlice...)
    return arr1
}