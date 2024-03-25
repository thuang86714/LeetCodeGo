package leetcodego

func findDuplicates(nums []int) []int {
    //credit to suman_buie, TC O(n), SC O(1)
    ans := []int{}
    var abs func(int)int
    abs = func(a int)int{
        if a < 0{
            a *= -1
        }
        return a
    }
    for _, num := range nums{
        if nums[abs(num) - 1] < 0{//this mean we have visited this num before
            ans = append(ans, abs(num))
        }
        nums[abs(num) - 1] = -1* nums[abs(num) - 1]
    }
    return ans
}