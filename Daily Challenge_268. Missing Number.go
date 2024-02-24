func missingNumber(nums []int) int {
    lens := len(nums)
    sum := lens*(lens + 1)/2
    for _, num := range(nums){
        sum -= num
    } 
    return sum
}