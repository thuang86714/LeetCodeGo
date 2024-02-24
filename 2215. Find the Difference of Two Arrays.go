func setDifference(a, b []int) []int{
    count := [2001]bool{}
    ans := []int{}
    for _,num := range(b){
        count[num + 1000] = true
    }
    for _,num := range(a){
        if !count[num + 1000]{
            ans = append(ans, num)
            count[num + 1000] = true
        }
    }
    return ans
}
func findDifference(nums1 []int, nums2 []int) [][]int {
    //credit to d9rs, TC O(n), SC O(1)
    return [][]int{setDifference(nums1, nums2), setDifference(nums2, nums1)}
}