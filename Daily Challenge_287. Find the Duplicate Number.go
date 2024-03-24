func findDuplicate(nums []int) int {
    left, right := 1, len(nums) - 1
    for left <= right{
        mid := left + (right - left)/2
        count := 0
        for _, num := range nums{
            if num <= mid{
                count++
            }
        }
        /*
        means the number of elements in the lower half 
        (1 to mid) is less than or equal to what it should be 
        in a scenario without duplicates. 
        Hence, the duplicate must be in the upper half, 
        so the search continues with left = mid + 1
        */
        if count <= mid{
            left = mid + 1
        }else{
            right = mid - 1
        }
    }
    /*
    the algorithm converges to the smallest number for which the count of numbers less than or equal to it is more than expected, which has to be the duplicate number.
    */
    return left
}