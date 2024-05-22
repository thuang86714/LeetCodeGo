package leetcodego

var sum int
func findEachSubsetSum(nums []int, idx, curXorVal int){
    if idx == len(nums){
        sum += curXorVal
        return
    }

    findEachSubsetSum(nums, idx + 1, curXorVal ^nums[idx])
    findEachSubsetSum(nums, idx + 1, curXorVal)
}
func subsetXORSum(nums []int) int {
    //TC O(2^n), SC O(n)
    sum = 0
    findEachSubsetSum(nums, 0, 0)
    return sum
}