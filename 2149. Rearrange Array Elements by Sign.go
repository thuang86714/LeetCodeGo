func rearrangeArray(nums []int) []int {
    //credi to kevinjunior, TC O(n), SC O(n)
    ans := make([]int, len(nums))
    pPtr, nPtr := 0, 1
    for _, num := range nums{
        if num > 0{
            ans[pPtr] = num
            pPtr += 2
        }else{
            ans[nPtr] = num
            nPtr += 2
        }
    }
    return ans
}