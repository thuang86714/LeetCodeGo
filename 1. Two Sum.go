package leetcodego

func twoSum(nums []int, target int) []int {
    //credit to krex0r
    dict := make(map[int]int, 0)
    for idx, num:= range nums{
        dict[num] = idx
    }

    for idx1, num := range nums{
        if idx2, ok := dict[target - num]; ok && idx1 != idx2{
            return []int{idx1, idx2}
        }
    }
    return []int{}
}