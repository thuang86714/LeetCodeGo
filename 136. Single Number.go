package leetcodego

func singleNumber(nums []int) int {
    //credit to satyamsinha93, TC O(n), SC O(1)
    result := 0
    for _, num := range nums{
        result ^= num
    }

    return result
}