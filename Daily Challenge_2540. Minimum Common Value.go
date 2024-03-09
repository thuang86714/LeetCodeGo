func getCommon(nums1 []int, nums2 []int) int {
    //TC O(n), SC O(1)
    idx1, idx2 := 0, 0
    for idx1 < len(nums1) && idx2 < len(nums2){
        if nums1[idx1] == nums2[idx2]{
            return nums2[idx2]
        }else if nums1[idx1] > nums2[idx2]{
            idx2++
        }else{
            idx1++
        }
    }
    return -1
}