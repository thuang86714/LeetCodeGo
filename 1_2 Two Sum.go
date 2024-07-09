package leetcodego

func twoSum4(nums []int, target int) []int {
    numMap := make(map[int]int)
    for idx, val := range nums{
        numMap[val] = idx
    }

    for idx, val := range nums{
        curTarget := target - val
        if curIdx, exist := numMap[curTarget]; exist && idx != curIdx{
            return []int{idx, curIdx}
        }
    }
    return []int{}
}